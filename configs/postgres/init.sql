CREATE TABLE IF NOT EXISTS wallets (
    id UUID PRIMARY KEY,
    balance DECIMAL(15, 2) NOT NULL DEFAULT 0.00 CHECK (balance >= 0)
);

INSERT INTO wallets (id) 
VALUES ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11') 
ON CONFLICT (id) DO NOTHING;