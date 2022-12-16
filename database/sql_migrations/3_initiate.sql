-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE datakesehatan (
    NIK VARCHAR(256) PRIMARY KEY,
    kelas VARCHAR(100),
    faskes VARCHAR(100),
    total_premi VARCHAR(256),
    no_bpjs VARCHAR(100),
    updated_at DATE,
    created_at DATE
)

-- +migrate StatementEnd