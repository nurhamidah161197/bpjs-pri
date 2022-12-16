package repository

import (
	"database/sql"
	"fmt"
	"pratice/structs"
	"time"
)

func GetDataPembayaranByPeriode(db *sql.DB, datapembayaran structs.DataPembayaran) (results []structs.DataPembayaran, err error) {
	sql := "SELECT * FROM datapembayaran WHERE nik=$1 AND periode=$2"
	rows, err := db.Query(sql, datapembayaran.NIK, datapembayaran.Periode)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&datapembayaran.Id, &datapembayaran.NIK, &datapembayaran.Periode, &datapembayaran.Premi, &datapembayaran.Updated_at)
		if err != nil {
			fmt.Println(err.Error())
		}

		results = append(results, datapembayaran)
	}

	return results, nil
}

func Bayar(db *sql.DB, bayar structs.DataPembayaran) (err error) {
	var datakesehatan structs.DataKesehatan
	var premi structs.Premi

	kelas := "SELECT kelas FROM datakesehatan WHERE nik=$1"
	errs := db.QueryRow(kelas, bayar.NIK).Scan(&datakesehatan.Kelas)
	_ = errs
	totalpremi := "SELECT premi FROM premi WHERE kelas=$1"
	errs = db.QueryRow(totalpremi, datakesehatan.Kelas).Scan(&premi.Premi)

	sql1 := "INSERT INTO datapembayaran (nik, periode, premi, created_at) VALUES ($1,$2,$3,$4)"
	errk := db.QueryRow(sql1, bayar.NIK, bayar.Periode, premi.Premi, time.Now())

	sql2 := "UPDATE datakesehatan SET total_premi=$1, updated_at=$2 WHERE nik=$3"
	errk = db.QueryRow(sql2, premi.Premi, time.Now(), bayar.NIK)

	if errk.Err() != nil {
		panic(errk.Err())
	}

	return errk.Err()
}
