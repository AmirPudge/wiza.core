CREATE TABLE accounts (
    id             UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    client_id      UUID         NOT NULL REFERENCES clients (id),
    account_number VARCHAR(34)  NOT NULL,
    currency       CHAR(3)      NOT NULL DEFAULT 'KZT',
    balance        NUMERIC(20,4) NOT NULL DEFAULT 0,
    type           TEXT         NOT NULL DEFAULT 'checking'
                       CHECK (type IN ('checking', 'savings', 'deposit')),
    status         TEXT         NOT NULL DEFAULT 'active'
                       CHECK (status IN ('active', 'blocked', 'closed')),
    opened_at      TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    closed_at      TIMESTAMPTZ,
    created_at     TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at     TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    deleted_at     TIMESTAMPTZ,

    CONSTRAINT accounts_number_unique  UNIQUE (account_number),
    CONSTRAINT accounts_balance_nonneg CHECK (balance >= 0),
    CONSTRAINT accounts_currency_upper CHECK (currency = UPPER(currency))
);

CREATE INDEX idx_accounts_client_id  ON accounts (client_id);
CREATE INDEX idx_accounts_status     ON accounts (status) WHERE deleted_at IS NULL;
CREATE INDEX idx_accounts_number     ON accounts (account_number);
