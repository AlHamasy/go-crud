package models

import (
	"database/sql"
	"time"

	"github.com/AlHamasy/go-crud/config"
	"github.com/AlHamasy/go-crud/entities"
)

type PasienModel struct {
	conn *sql.DB
}

func NewPasienModel() *PasienModel {
	conn, err := config.DBConnection()
	if err != nil {
		panic(err)
	}

	return &PasienModel{
		conn: conn,
	}
}

func (p *PasienModel) FindAll() ([]entities.Pasien, error) {

	rows, err := p.conn.Query("select * from pasien")
	if err != nil {
		return []entities.Pasien{}, err
	}
	defer rows.Close()

	var dataPasien []entities.Pasien
	for rows.Next() {
		var pasien entities.Pasien
		rows.Scan(&pasien.Id, &pasien.NamaLengkap, &pasien.NIK, &pasien.JenisKelamin, &pasien.TempatLahir, &pasien.TanggalLahir, &pasien.Alamat, &pasien.NoHp)

		if pasien.JenisKelamin == "1" {
			pasien.JenisKelamin = "Laki-laki"
		} else if pasien.JenisKelamin == "2" {
			pasien.JenisKelamin = "Perempuan"
		}

		// 2006-01-02 -> yyyy-mm-dd
		tglLahir, _ := time.Parse("2006-01-02", pasien.TanggalLahir)
		// 02 January 2006 -> dd MMMM yyyy
		pasien.TanggalLahir = tglLahir.Format("02 January 2006")

		dataPasien = append(dataPasien, pasien)
	}

	return dataPasien, nil
}

func (p *PasienModel) Create(pasien entities.Pasien) (bool, error) {

	result, err := p.conn.Exec("insert into pasien (nama_lengkap, nik, jenis_kelamin, tempat_lahir, tanggal_lahir, alamat, no_hp) values (?,?,?,?,?,?,?)",
		pasien.NamaLengkap, pasien.NIK, pasien.JenisKelamin, pasien.TempatLahir, pasien.TanggalLahir, pasien.Alamat, pasien.NoHp)

	if err != nil {
		return false, err
	}
	lastInsertId, _ := result.LastInsertId()
	return lastInsertId > 0, nil
}

func (p *PasienModel) Find(id int64, pasien *entities.Pasien) error {

	return p.conn.QueryRow("select * from pasien where id = ?", id).Scan(
		&pasien.Id,
		&pasien.NamaLengkap,
		&pasien.NIK,
		&pasien.JenisKelamin,
		&pasien.TempatLahir,
		&pasien.TanggalLahir,
		&pasien.Alamat,
		&pasien.NoHp,
	)

}

func (p *PasienModel) Update(pasien entities.Pasien) error {

	_, err := p.conn.Exec(
		"update pasien set nama_lengkap = ?, nik = ?, jenis_kelamin = ?, tempat_lahir = ?, tanggal_lahir = ?, alamat = ?, no_hp = ? where id = ?",
		pasien.NamaLengkap, pasien.NIK, pasien.JenisKelamin, pasien.TempatLahir, pasien.TanggalLahir, pasien.Alamat, pasien.NoHp, pasien.Id)

	if err != nil {
		return err
	}

	return nil
}

func (p *PasienModel) Delete(id int64) {
	p.conn.Exec("delete from pasien where id = ?", id)
}
