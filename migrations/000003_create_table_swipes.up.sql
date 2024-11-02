CREATE TABLE
    IF NOT EXISTS swipes (
        id VARCHAR(255) PRIMARY KEY,
        swiper_id VARCHAR(255) REFERENCES users (id) ON DELETE CASCADE,
        swiped_id VARCHAR(255) REFERENCES users (id) ON DELETE CASCADE,
        swipe_type VARCHAR(10),
        swiped_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );