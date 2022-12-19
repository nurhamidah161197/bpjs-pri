package repository

import (
	"database/sql"
	"pratice/structs"
	"time"
)

func GetAllMaster(db *sql.DB) (results []structs.Master, err error) {
	sql := "SELECT * FROM masterdata"

	rows, err := db.Query(sql)

	if err != nil {

	}
	defer rows.Close()
	for rows.Next() {
		var master = structs.Master{}
		err = rows.Scan(&master.Nama, &master.NIK, &master.Email, &master.Gender, &master.TglLahir,
			&master.NoHp, &master.Alamat, &master.Created_at, &master.Updated_at)
		if err != nil {
			return nil, err
		}

		results = append(results, master)
	}

	return
}

func InsertMaster(db *sql.DB, master structs.Master) (err error) {
	sql1 := "INSERT INTO masterdata (nik, nama, email, gender, tgl_lahir, no_hp, alamat, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)"
	errs := db.QueryRow(sql1, master.NIK, master.Nama, master.Email, master.Gender, master.TglLahir, master.NoHp, master.Alamat, time.Now(), time.Now())

	if errs.Err() != nil {
		panic(errs.Err())
	}

	return errs.Err()
}

func UpdateMaster(db *sql.DB, master structs.Master) (err error) {
	sql := "UPDATE masterdata SET nama=$1, email=$2, gender=$3, tgl_lahir=$4, no_hp=$5,  alamat=$6, updated_at=$7 WHERE nik=$8"
	errs := db.QueryRow(sql, master.Nama, master.Email, master.Gender, master.TglLahir, master.Alamat, master.NoHp, time.Now(), master.NIK)
	return errs.Err()
}

func DeleteMaster(db *sql.DB, master structs.Master) (err error) {
	sql := "DELETE FROM masterdata WHERE nik = $1"
	errs := db.QueryRow(sql, master.NIK)

	if errs != nil {
		panic(errs.Err())
	}

	if errs != nil {
		panic(errs.Err())
	}

	return errs.Err()
}
