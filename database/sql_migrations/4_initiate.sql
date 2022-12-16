-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE datapembayaran (
    id SERIAL PRIMARY KEY,
    NIK VARCHAR(256),
    periode VARCHAR(50),
    premi VARCHAR(100),
    created_at DATE
)
-- +migrate StatementEnd