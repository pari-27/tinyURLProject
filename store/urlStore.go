package store

import (
	"database/sql"
	"fmt"
	"github.com/pari-27/tinyURLProject/utils"
)

type Store interface {
	Create(utils.URL) (err error)
	Get(string) (utils.URL, error)
	Delete(utils.URL) (err error)
}

func Init() (db *sql.DB, err error) {
	db, err = sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/test")
	return
}

type TinyUrl struct {
	DB *sql.DB
}

var urlDBStore = map[string]string{}

func (tu *TinyUrl) Create(u utils.URL) (err error) {
	//query := fmt.Sprintf("INSERT INTO urls VALUES (%s,%s)", u.EncodedURL, u.LongURL)
	//_, err = tu.DB.Query(query)
	//if err != nil {
	//	return
	//}

	if _, ok := urlDBStore[u.EncodedURL]; !ok {
		urlDBStore[u.EncodedURL] = u.LongURL
		fmt.Println("key added")
	}
	fmt.Println("key not added")

	return fmt.Errorf("key already exists")
}
func (tu *TinyUrl) Delete(u utils.URL) (err error) {
	//query := fmt.Sprintf("DELETE * FROM urls WHERE encodedURL = %s )", u.EncodedURL)
	//_, err = tu.DB.Query(query)
	//if err != nil {
	//	return
	//}

	if _, ok := urlDBStore[u.EncodedURL]; !ok {
		return fmt.Errorf("key not found")
	}
	return

}

func (tu *TinyUrl) Get(string) (u utils.URL, err error) {
	//query := fmt.Sprintf("SELECT * FROM urls WHERE encodedURL = %s", u.EncodedURL)
	//err = tu.DB.QueryRow(query).Scan(&u.LongURL)
	//if err != nil {
	//	return
	//}
	//return
	if _, ok := urlDBStore[u.EncodedURL]; !ok {
		return utils.URL{}, fmt.Errorf("key not found")
	}
	u.LongURL = urlDBStore[u.EncodedURL]
	return
}
