CREATE TABLE IF NOT EXISTS staffs (
    staff_id VARCHAR(36) NOT NULL,
    position_id VARCHAR(36) NOT NULL,
    first_name VARCHAR(150) NOT NULL,
    last_name VARCHAR(150) NOT NULL,
    email VARCHAR(150) NOT NULL,
    password_hash BYTEA NOT NULL,
    password_salt BYTEA NOT NULL,
    status VARCHAR(40) NOT NULL,
    selfie_url VARCHAR(200),
    selfie_thumb_url VARCHAR(200),
    updated_at TIMESTAMP,
    created_at TIMESTAMP NOT NULL,
    CONSTRAINT pk_staffs PRIMARY KEY (staff_id),
    CONSTRAINT fk_staffs_position
        FOREIGN KEY (position_id)
        REFERENCES positions(position_id)
);