CREATE TABLE IF NOT EXISTS accounts (
    id VARCHAR(36) PRIMARY KEY,
    balance DECIMAL(12,2) NOT NULL DEFAULT 0,
    is_enabled TINYINT(1) UNSIGNED NOT NULL DEFAULT 0,
    -- token VARCHAR(50) NOT NULL,
    created_at timestamp NOT NULL DEFAULT NOW(),
    updated_at timestamp NOT NULL DEFAULT NOW(),
    deleted_at timestamp,
    INDEX(id, is_enabled, created_at, deleted_at)
) COMMENT 'Master table for accounts';

CREATE TABLE IF NOT EXISTS transactions (
    id VARCHAR(36) PRIMARY KEY,
    account_id VARCHAR(36) NOT NULL,
    reference_id VARCHAR(36) NOT NULL,
    type ENUM('deposit', 'withdrawal') NOT NULL,
    amount DECIMAL(12, 2) NOT NULL,
    status ENUM('success', 'failed') NOT NULL,
    created_at timestamp NOT NULL DEFAULT NOW(),
    updated_at timestamp NOT NULL DEFAULT NOW(),
    deleted_at timestamp,
    INDEX(id, reference_id, created_at, deleted_at),
    CONSTRAINT fk_transactions_accounts FOREIGN KEY (account_id) REFERENCES accounts(id) ON UPDATE CASCADE ON DELETE CASCADE
) COMMENT 'Transaction table for money movements';
