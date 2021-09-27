package biodata_diri

import (
	"api-mysql/config"
	"api-mysql/models"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

const (
	table          = "biodata_diri"
	layoutDateTime = "2006-01-02 15:04:05"
)

// GetAll Biodata_diri
func GetAll(ctx context.Context) ([]models.Biodata_diri, error) {

	var Biodata_diris []models.Biodata_diri

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}

	queryText := fmt.Sprintf("SELECT * FROM %v Order By Biodata_diri_id DESC", table)

	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var Biodata_diri models.Biodata_diri

		if err = rowQuery.Scan(&Biodata_diri.Biodata_id,
			&Biodata_diri.Nama,
			&Biodata_diri.Tanggal_lahir,
			&Biodata_diri.Jenis_kelamin,
			&Biodata_diri.Email,
			&Biodata_diri.No_hp,
			&Biodata_diri.Profil_pic,
			&Biodata_diri.Users_id); err != nil {
			return nil, err
		}

		Biodata_diris = append(Biodata_diris, Biodata_diri)
	}

	return Biodata_diris, nil
}

// Insert Biodata_diri
func Insert(ctx context.Context, Biodata_diri models.Biodata_diri) error {
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("INSERT INTO %v (biodata_id, nama, tanggal_lahir, jenis_kelamin, email, no_hp, profil_pic, users_id) values('%v','%v','%v','%v','%v','%v','%v','%v')", table,
		&Biodata_diri.Biodata_id,
		&Biodata_diri.Nama,
		&Biodata_diri.Tanggal_lahir,
		&Biodata_diri.Jenis_kelamin,
		&Biodata_diri.Email,
		&Biodata_diri.No_hp,
		&Biodata_diri.Profil_pic,
		&Biodata_diri.Users_id,
	)

	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}
	return nil
}

// Update Biodata_diri
func Update(ctx context.Context, Biodata_diri models.Biodata_diri, id string) error {

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("UPDATE %v set biodata_diri_id ='%d', nama ='%s', tanggal_lahir ='%s', jenis_kelamin ='%s', email ='%s', no_hp ='%d', profil_pic ='%s', users_id ='%d' where biodata_diri_id = %s",
		table,
		Biodata_diri.Biodata_id,
		Biodata_diri.Nama,
		Biodata_diri.Tanggal_lahir,
		Biodata_diri.Jenis_kelamin,
		Biodata_diri.Email,
		Biodata_diri.No_hp,
		Biodata_diri.Profil_pic,
		Biodata_diri.Users_id,
		id,
	)

	fmt.Println(queryText)

	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}

	return nil
}

// Delete Biodata_diri
func Delete(ctx context.Context, id string) error {
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("DELETE FROM %v where biodata_diri_id = %s", table, id)

	s, err := db.ExecContext(ctx, queryText)

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	check, err := s.RowsAffected()
	fmt.Println(check)
	if check == 0 {
		return errors.New("id tidak ada")
	}

	if err != nil {
		fmt.Println(err.Error())
	}

	return nil
}
