package store

import (
	"../utils"
	"database/sql"
	"fmt"
)

type Store interface {
	Create(utils.URL) (err error)
	Get() (utils.URL, error)
	Delete(utils.URL) (err error)
}

func Init() (db *sql.DB, err error) {
	db, err = sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/test")
	return
}

type TinyUrl struct {
	DB *sql.DB
}

func (tu *TinyUrl) Create(u utils.URL) (err error) {
	query := fmt.Sprintf("INSERT INTO urls VALUES (%s,%s)", u.EncodedURL, u.LongURL)
	_, err = tu.DB.Query(query)
	if err != nil {
		return
	}
	return
}
func (tu *TinyUrl) Delete(u utils.URL) (err error) {
	query := fmt.Sprintf("DELETE * FROM urls WHERE encodedURL = %s )", u.EncodedURL)
	_, err = tu.DB.Query(query)
	if err != nil {
		return
	}
	return
}

func (tu *TinyUrl) Get(u utils.URL) (err error) {
	query := fmt.Sprintf("SELECT * FROM urls WHERE encodedURL = %s", u.EncodedURL)
	err = tu.DB.QueryRow(query).Scan(&u)
	if err != nil {
		return
	}
	return
}
