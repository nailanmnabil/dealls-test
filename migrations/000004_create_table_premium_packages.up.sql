CREATE TABLE
    IF NOT EXISTS premium_packages (
        id VARCHAR(255) PRIMARY KEY,
        package_name VARCHAR(100) NOT NULL,
        description TEXT,
        price NUMERIC(10, 2) NOT NULL,
        feature_type VARCHAR(255),
        active_period VARCHAR(255),
        created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
    );