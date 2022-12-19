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
		err = rows.Scan(&datakesehatan.NIK, &datakesehatan.Kelas, &datakesehatan.Faskes, &datakesehatan.TotalPremi, &datakesehatan.NoBPJS, &datakesehatan.Updated_at, &datakesehatan.Created_at)
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

func UpdateDataKesehatan(db *sql.DB, datakesehatan structs.DataKesehatan) (err error) {
	sql := "UPDATE datakesehatan SET faskes=$1, updated_at=$2, no_bpjs=$3 WHERE nik=$4"
	errs := db.QueryRow(sql, datakesehatan.Faskes, time.Now(), datakesehatan.NoBPJS, datakesehatan.NIK)
	return errs.Err()
}

func Tagihan(db *sql.DB, datakesehatan structs.DataKesehatan) (results structs.DataKesehatan, total_tagihan int64) {
	var datapembayaran structs.DataPembayaran
	var premi structs.Premi

	sql := "SELECT * FROM datakesehatan WHERE nik=$1"
	db.QueryRow(sql, datakesehatan.NIK).Scan(&datakesehatan.NIK, &datakesehatan.Kelas, &datakesehatan.Faskes, &datakesehatan.NoBPJS, &datakesehatan.TotalPremi, &datakesehatan.Updated_at, &datakesehatan.Created_at)

	sql2 := "SELECT periode FROM datapembayaran WHERE nik=$1 ORDER BY id DESC"
	db.QueryRow(sql2, datakesehatan.NIK).Scan(&datapembayaran.Periode)

	sql3 := "SELECT premi FROM premi WHERE kelas=$1"
	db.QueryRow(sql3, datakesehatan.Kelas).Scan(&premi.Premi)

	timePremi, err := time.Parse("2006-01-02", datapembayaran.Periode+"-01")
	if err != nil {
		panic(err)
	}
	today := time.Now()
	difference := today.Sub(timePremi)
	countDifference := int64(difference.Hours() / 24 / 30)

	total_tagihan = countDifference * premi.Premi
	return datakesehatan, total_tagihan
}

func DeleteDataKesehatan(db *sql.DB, datakesehatan structs.DataKesehatan) (err error) {
	sql := "DELETE FROM datakesehatan WHERE nik = $1"
	errs := db.QueryRow(sql, datakesehatan.NIK)

	if errs != nil {
		panic(errs.Err())
	}

	return errs.Err()
}
