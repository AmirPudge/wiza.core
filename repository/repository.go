package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"wiza.core/domain"
)

type clientRepository struct {
	db *pgxpool.Pool
}

func NewClientRepository(db *pgxpool.Pool) domain.ClientRepository {
	return &clientRepository{db: db}
}

func (r *clientRepository) GetByIIN(ctx context.Context, iin string) (*domain.Client, error) {
	client, err := r.fetchClient(ctx, iin)
	if err != nil {
		return nil, err
	}

	accounts, err := r.fetchAccounts(ctx, client.ID)
	if err != nil {
		return nil, err
	}

	if len(accounts) > 0 {
		accountIDs := make([]string, len(accounts))
		for i, a := range accounts {
			accountIDs[i] = a.ID
		}

		cards, err := r.fetchCards(ctx, accountIDs)
		if err != nil {
			return nil, err
		}

		cardsByAccount := make(map[string][]domain.Card)
		for _, c := range cards {
			cardsByAccount[c.AccountID] = append(cardsByAccount[c.AccountID], c)
		}

		for i := range accounts {
			accounts[i].Cards = cardsByAccount[accounts[i].ID]
		}
	}

	client.Accounts = accounts
	return client, nil
}

func (r *clientRepository) fetchClient(ctx context.Context, iin string) (*domain.Client, error) {
	const q = `
		SELECT id, iin, first_name, last_name, middle_name,
		       birth_date, phone, email, status
		FROM clients
		WHERE iin = $1 AND deleted_at IS NULL`

	var c domain.Client
	err := r.db.QueryRow(ctx, q, iin).Scan(
		&c.ID, &c.IIN, &c.FirstName, &c.LastName, &c.MiddleName,
		&c.BirthDate, &c.Phone, &c.Email, &c.Status,
	)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, domain.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *clientRepository) fetchAccounts(ctx context.Context, clientID string) ([]domain.Account, error) {
	const q = `
		SELECT id, client_id, account_number, currency,
		       balance::float8, type, status, opened_at
		FROM accounts
		WHERE client_id = $1 AND deleted_at IS NULL
		ORDER BY opened_at`

	rows, err := r.db.Query(ctx, q, clientID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var accounts []domain.Account
	for rows.Next() {
		var a domain.Account
		if err := rows.Scan(
			&a.ID, &a.ClientID, &a.AccountNumber, &a.Currency,
			&a.Balance, &a.Type, &a.Status, &a.OpenedAt,
		); err != nil {
			return nil, err
		}
		accounts = append(accounts, a)
	}
	return accounts, rows.Err()
}

func (r *clientRepository) fetchCards(ctx context.Context, accountIDs []string) ([]domain.Card, error) {
	const q = `
		SELECT id, account_id, masked_pan, cardholder_name,
		       expiry_month, expiry_year, payment_system, type, status, issued_at
		FROM cards
		WHERE account_id = ANY($1) AND deleted_at IS NULL
		ORDER BY issued_at`

	rows, err := r.db.Query(ctx, q, accountIDs)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cards []domain.Card
	for rows.Next() {
		var c domain.Card
		if err := rows.Scan(
			&c.ID, &c.AccountID, &c.MaskedPAN, &c.CardholderName,
			&c.ExpiryMonth, &c.ExpiryYear, &c.PaymentSystem, &c.Type, &c.Status, &c.IssuedAt,
		); err != nil {
			return nil, err
		}
		cards = append(cards, c)
	}
	return cards, rows.Err()
}
