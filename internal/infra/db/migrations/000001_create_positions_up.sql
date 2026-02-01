CREATE TABLE IF NOT EXISTS positions (
    position_id VARCHAR(36) NOT NULL,
    name VARCHAR(200) NOT NULL,
    updated_at TIMESTAMP,
    created_at TIMESTAMP NOT NULL,
    CONSTRAINT pk_positions PRIMARY KEY (position_id)
);