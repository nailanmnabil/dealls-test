CREATE TABLE
    IF NOT EXISTS profile_visits (
        id VARCHAR(255) PRIMARY KEY,
        visitor_id VARCHAR(255) REFERENCES users (id) ON DELETE CASCADE,
        visited_id VARCHAR(255) REFERENCES users (id) ON DELETE CASCADE,
        visit_date DATE DEFAULT CURRENT_DATE,
        swiped_at TIMESTAMP DEFAULT NULL,
        UNIQUE (visitor_id, visited_id, visit_date)
    );