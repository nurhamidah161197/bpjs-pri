package repository

import (
	"database/sql"
	"pratice/structs"
	"time"
)

func GetAllDataKesehatan(db *sql.DB) (err error, results []structs.DataKesehatan) {
	sql := "SELECT * FROM datakesehatan"

	rows, err := db.Query(sql)

	if err != nil {

	}
	defer rows.Close()
	for rows.Next() {
		var datakesehatan = structs.DataKesehatan{}
		err = rows.Scan(&datakesehatan.NIK, &datakesehatan.Kelas, &datakesehatan.Faskes, &datakesehatan.NoBPJS, &datakesehatan.TotalPremi, &datakesehatan.Updated_at, &datakesehatan.Created_at)
		if err != nil {

		}

		results = append(results, datakesehatan)
	}

	return
}

func GetDataKesehatanByNIK(db *sql.DB, datakesehatan structs.DataKesehatan) (results []structs.DataKesehatan, err error) {
	sql := "SELECT * FROM datakesehatan WHERE nik=$1"

	rows, err := db.Query(sql, datakesehatan.NIK)

	if err != nil {

	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&datakesehatan.NIK, &datakesehatan.Kelas, &datakesehatan.Faskes, &datakesehatan.NoBPJS, &datakesehatan.TotalPremi, &datakesehatan.Updated_at, &datakesehatan.Created_at)
		if err != nil {

		}

		results = append(results, datakesehatan)
	}

	return results, nil
}

func InsertDataKesehatan(db *sql.DB, datakesehatan structs.DataKesehatan) (err error) {
	total_premi := "0"
	sql := "INSERT INTO datakesehatan (nik , kelas, faskes,total_premi, created_at, updated_at, no_bpjs) VALUES ($1, $2, $3, $4, $5, $6, $7)"
	errs := db.QueryRow(sql, datakesehatan.NIK, datakesehatan.Kelas, datakesehatan.Faskes, total_premi, time.Now(), time.Now(), datakesehatan.NoBPJS)

	if errs.Err() != nil {
		panic(errs.Err())
	}

	return errs.Err()
}

func UpdateFaskes(db *sql.DB, datakesehatan structs.DataKesehatan) (err error) {
	sql := "UPDATE datakesehatan SET faskes=$1, updated_at=$2 WHERE nik=$3"
	errs := db.QueryRow(sql, datakesehatan.Faskes, time.Now(), datakesehatan.NIK)
	return errs.Err()
}

func Tagihan(db *sql.DB, datakesehatan structs.DataKesehatan) (results structs.DataKesehatan, total_tagihan int64) {
	var datapembayaran structs.DataPembayaran
	var premi structs.Premi

	sql := "SELECT kelas FROM datakesehatan WHERE datakesehatan WHERE nik=$1"
	db.QueryRow(sql, datakesehatan.NIK).Scan(&datakesehatan.Kelas)

	sql2 := "SELECT periode FROM datapembayaran WHERE nik=$1 ORDER BY id DESC"
	db.QueryRow(sql2, datakesehatan.NIK).Scan(&datapembayaran.Periode)

	sql3 := "SELECT premi FROM premi WHERE kelas=$1"
	db.QueryRow(sql3, datakesehatan.Kelas).Scan(&premi.Premi)

	total_tagihan = 1 * premi.Premi

	return datakesehatan, total_tagihan
}
