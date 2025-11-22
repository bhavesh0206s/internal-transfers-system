CREATE TABLE IF NOT EXISTS accounts (
	account_id BIGINT PRIMARY KEY,
	balance BIGINT NOT NULL DEFAULT 0
);

CREATE TABLE IF NOT EXISTS transactions (
	id BIGSERIAL PRIMARY KEY,
	source_account_id BIGINT REFERENCES accounts(account_id),
  destination_account_id BIGINT REFERENCES accounts(account_id),
	amount BIGINT NOT NULL DEFAULT 0,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_source_account_id ON transactions(source_account_id);
CREATE INDEX IF NOT EXISTS idx_destination_account_id ON transactions(destination_account_id);