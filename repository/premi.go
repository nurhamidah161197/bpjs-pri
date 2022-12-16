package repository

import (
	"database/sql"
	"pratice/structs"
	"time"
)

func GetAllPremi(db *sql.DB) (err error, results []structs.Premi) {
	sql := "SELECT * FROM premi"

	rows, err := db.Query(sql)

	if err != nil {

	}
	defer rows.Close()
	for rows.Next() {
		var premi = structs.Premi{}
		err = rows.Scan(&premi.Id, &premi.Kelas, &premi.Premi, &premi.Created_at, &premi.Updated_at)
		if err != nil {

		}

		results = append(results, premi)
	}

	return
}

func InsertPremi(db *sql.DB, premi structs.Premi) (err error) {
	sql := "INSERT INTO premi (kelas, premi, updated_at, created_at) VALUES ($1,$2,$3,$4)"
	errs := db.QueryRow(sql, premi.Kelas, premi.Premi, time.Now(), time.Now())
	return errs.Err()
}

func UpdatePremi(db *sql.DB, premi structs.Premi) (err error) {

	sql := "UPDATE premi SET kelas=$1, premi =$2, updated_at=$3 WHERE id=$4"
	errs := db.QueryRow(sql, premi.Kelas, premi.Premi, time.Now(), premi.Id)
	return errs.Err()
}

func DeletePremi(db *sql.DB, premi structs.Premi) (err error) {
	sql := "DELETE FROM premi WHERE id = $1"
	errs := db.QueryRow(sql, premi.Id)

	return errs.Err()
}
