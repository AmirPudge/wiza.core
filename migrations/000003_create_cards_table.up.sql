CREATE TABLE cards (
    id               UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    account_id       UUID         NOT NULL REFERENCES accounts (id),
    masked_pan       VARCHAR(19)  NOT NULL,
    cardholder_name  VARCHAR(255) NOT NULL,
    expiry_month     SMALLINT     NOT NULL CHECK (expiry_month BETWEEN 1 AND 12),
    expiry_year      SMALLINT     NOT NULL,
    payment_system   TEXT         NOT NULL
                         CHECK (payment_system IN ('visa', 'mastercard', 'mir', 'unionpay')),
    type             TEXT         NOT NULL DEFAULT 'debit'
                         CHECK (type IN ('debit', 'credit', 'virtual', 'prepaid')),
    status           TEXT         NOT NULL DEFAULT 'active'
                         CHECK (status IN ('active', 'blocked', 'expired', 'closed')),
    issued_at        TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    created_at       TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at       TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    deleted_at       TIMESTAMPTZ,

    CONSTRAINT cards_expiry_year_sane CHECK (expiry_year BETWEEN 2020 AND 2099)
);

CREATE INDEX idx_cards_account_id ON cards (account_id);
CREATE INDEX idx_cards_status     ON cards (status) WHERE deleted_at IS NULL;
CREATE INDEX idx_cards_masked_pan ON cards (masked_pan);
