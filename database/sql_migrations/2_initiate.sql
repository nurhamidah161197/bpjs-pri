-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE masterdata (
    nama VARCHAR(256),
    nik VARCHAR(256) PRIMARY KEY,
    email VARCHAR(256),
    gender VARCHAR(256),
    tgl_lahir DATE,
    alamat VARCHAR(256),
    no_hp VARCHAR(256),
    created_at DATE,
    updated_at DATE
)

-- +migrate StatementEnd