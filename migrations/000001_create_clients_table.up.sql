CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE clients (
    id          UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    iin         CHAR(12)    NOT NULL,
    first_name  VARCHAR(100) NOT NULL,
    last_name   VARCHAR(100) NOT NULL,
    middle_name VARCHAR(100),
    birth_date  DATE         NOT NULL,
    phone       VARCHAR(20),
    email       VARCHAR(255),
    status      TEXT         NOT NULL DEFAULT 'active'
                    CHECK (status IN ('active', 'blocked', 'closed')),
    created_at  TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    deleted_at  TIMESTAMPTZ,

    CONSTRAINT clients_iin_unique UNIQUE (iin),
    CONSTRAINT clients_iin_format CHECK (iin ~ '^[0-9]{12}$')
);

CREATE INDEX idx_clients_iin        ON clients (iin);
CREATE INDEX idx_clients_status     ON clients (status) WHERE deleted_at IS NULL;
CREATE INDEX idx_clients_deleted_at ON clients (deleted_at);
