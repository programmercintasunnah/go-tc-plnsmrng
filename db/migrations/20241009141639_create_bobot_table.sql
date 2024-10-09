-- migrate:up
CREATE TABLE IF NOT EXISTS bobot (
    id SERIAL PRIMARY KEY,
    parent_id INT,
    nama VARCHAR(255) NOT NULL,
    nomor VARCHAR(50) NOT NULL
);

-- migrate:down
DROP TABLE IF EXISTS bobot;