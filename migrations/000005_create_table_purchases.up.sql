CREATE TABLE
    IF NOT EXISTS purchases (
        id VARCHAR(255) PRIMARY KEY,
        user_id VARCHAR(255) REFERENCES users (id) ON DELETE CASCADE,
        package_id VARCHAR(255) REFERENCES premium_packages (id) ON DELETE CASCADE,
        purchase_date TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
        expired_date TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
    );