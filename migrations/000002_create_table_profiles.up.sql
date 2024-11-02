CREATE TABLE
    IF NOT EXISTS profiles (
        id VARCHAR(255) PRIMARY KEY,
        user_id VARCHAR(255) UNIQUE REFERENCES users (id) ON DELETE CASCADE,
        bio TEXT,
        age INT,
        location VARCHAR(100),
        profile_pic_url VARCHAR(255),
        created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
    );