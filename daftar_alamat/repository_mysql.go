package daftar_alamat

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
	table          = "daftar_alamat"
	layoutDateTime = "2006-01-02 15:04:05"
)

// GetAll Daftar_alamat
func GetAll(ctx context.Context) ([]models.Daftar_alamat, error) {

	var Daftar_alamats []models.Daftar_alamat

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}

	queryText := fmt.Sprintf("SELECT * FROM %v Order By alamat_id DESC", table)

	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var Daftar_alamat models.Daftar_alamat

		if err = rowQuery.Scan(&Daftar_alamat.Alamat_id,
			&Daftar_alamat.Label_alamat,
			&Daftar_alamat.Nama_penerima,
			&Daftar_alamat.No_hp,
			&Daftar_alamat.Kota_kecamatan,
			&Daftar_alamat.Kode_pos,
			&Daftar_alamat.Alamat,
			&Daftar_alamat.Users_id); err != nil {
			return nil, err
		}

		Daftar_alamats = append(Daftar_alamats, Daftar_alamat)
	}

	return Daftar_alamats, nil
}

// Insert Daftar_alamat
func Insert(ctx context.Context, Daftar_alamat models.Daftar_alamat) error {
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("INSERT INTO %v (alamat_id, label_alamat, nama_penerima, no_hp, kota_kecamatan, kode_pos, alamat, users_id) values('%v','%v','%v','%v','%v','%v','%v','%v')", table,
		Daftar_alamat.Alamat_id,
		Daftar_alamat.Label_alamat,
		Daftar_alamat.Nama_penerima,
		Daftar_alamat.No_hp,
		Daftar_alamat.Kota_kecamatan,
		Daftar_alamat.Kode_pos,
		Daftar_alamat.Alamat,
		Daftar_alamat.Users_id,
	)

	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}
	return nil
}

// Update Daftar_alamat
func Update(ctx context.Context, Daftar_alamat models.Daftar_alamat, id string) error {

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("UPDATE %v set alamat_id ='%d', label_alamat ='%s', nama_penerima ='%s', no_hp ='%d', kota_kecamatan ='%s', kode_pos ='%d', alamat ='%s', users_id ='%d' where alamat_id = %s",
		table,
		Daftar_alamat.Alamat_id,
		Daftar_alamat.Label_alamat,
		Daftar_alamat.Nama_penerima,
		Daftar_alamat.No_hp,
		Daftar_alamat.Kota_kecamatan,
		Daftar_alamat.Kode_pos,
		Daftar_alamat.Alamat,
		Daftar_alamat.Users_id,
		id,
	)

	fmt.Println(queryText)

	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}

	return nil
}

// Delete Daftar_alamat
func Delete(ctx context.Context, id string) error {
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("DELETE FROM %v where alamat_id = %s", table, id)

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
