package saver

import (
	"database/sql"
	"github.com/ProSt1ll/UrlCutterAPI/model"
	"log"
	//_ "github.com/golang-migrate/migrate/v4/database/postgres"
	//_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"net/url"
)

type DBSaver struct {
	DB     *sql.DB
	Host   string
	Port   string
	DBName string
}

func NewDBSaver(db *sql.DB) Saver {
	return &DBSaver{
		DB: db,
	}
}

func NewDB(host string, port string, dbname string) DBSaver {
	b := DBSaver{
		Host:   host,
		Port:   port,
		DBName: dbname,
	}
	if err := b.Open(); err != nil {
		panic(err)
	}
	return b
}

// Open connection to database
func (b *DBSaver) Open() error {
	db, err := sql.Open("postgres", "host="+b.Host+" port="+b.Port+" user=postgres password=medusa dbname="+b.DBName+" sslmode=disable ")
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	b.DB = db
	return nil
}

func (b *DBSaver) Close() error {
	err := b.DB.Close()
	if err != nil {
		return err
	}
	return nil
}

// StoreURL stores model of
func (b *DBSaver) StoreURL(urls model.URLs) (int, error) {

	if err := b.DB.QueryRow("INSERT INTO urls (long_url,short_url) VALUES ($1,$2) RETURNING id",
		urls.LongUrl.String(),
		urls.ShortUrl,
	).Scan(&urls.Id); err != nil {
		return 0, err
	}
	return urls.Id, nil
}

// LoadShort load's short URL with full
func (b *DBSaver) LoadShort(key url.URL) (model.URLs, bool) {

	u := model.URLs{}
	var temp string
	if err := b.DB.QueryRow("SELECT id, long_url, short_url FROM urls WHERE long_url  = $1", key.String()).Scan(
		&u.Id,
		&temp,
		&u.ShortUrl); err != nil {
		return u, false
	}

	o, err := url.Parse(temp)
	if err != nil {
		panic(err)
	}

	u.LongUrl = *o
	return u, true
}

// LoadLong load's full URL with short
func (b *DBSaver) LoadLong(key string) (model.URLs, bool) {
	u := model.URLs{}
	var temp string
	if err := b.DB.QueryRow("SELECT id, long_url, short_url FROM URLs WHERE short_url  = $1", key).Scan(
		&u.Id,
		&temp,
		&u.ShortUrl); err != nil {
		return u, false
	}
	o, _ := url.Parse(temp)
	u.LongUrl = *o
	return u, true
}
