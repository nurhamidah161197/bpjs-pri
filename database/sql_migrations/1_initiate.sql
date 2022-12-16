-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE premi (
    id SERIAL PRIMARY KEY,
    kelas VARCHAR(256),
    premi INTEGER,
    updated_at DATE,
    created_at DATE
)

-- +migrate StatementEnd