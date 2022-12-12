package saver

import (
	"database/sql"
	"github.com/ProSt1ll/UrlCutterAPI/model"
	_ "github.com/lib/pq"
	"net/url"
)

type DBSaver struct {
	db     *sql.DB
	host   string
	port   string
	dbname string
}

func NewDBSaver(db *sql.DB) Saver {
	return &DBSaver{
		db: db,
	}
}

func NewDB() Saver {
	return &DBSaver{}
}

func (b *DBSaver) Config(host string, port string, dbname string) {
	b.host = host
	b.port = port
	b.dbname = dbname
}

func (b *DBSaver) Open() error {
	db, err := sql.Open("postgres", "host="+b.host+" port="+b.port+" user=postgres password=medusa dbname="+b.dbname+" sslmode=disable ")
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	b.db = db

	return nil
}

func (b *DBSaver) Close() error {
	err := b.db.Close()
	if err != nil {
		return err
	}
	return nil
}

func (b *DBSaver) StoreURL(urls model.URLs) (int, error) {
	if err := b.db.QueryRow("INSERT INTO urls (long_url,short_url) VALUES ($1,$2) RETURNING id",
		urls.LongUrl.String(),
		urls.ShortUrl,
	).Scan(&urls.Id); err != nil {
		return 0, err
	}
	return urls.Id, nil
}

func (b *DBSaver) LoadShort(key url.URL) (model.URLs, bool) {
	u := model.URLs{}
	var temp string
	if err := b.db.QueryRow("SELECT id, long_url, short_url FROM urls WHERE long_url  = $1", key.String()).Scan(
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

func (b *DBSaver) LoadLong(key string) (model.URLs, bool) {
	u := model.URLs{}
	var temp string
	if err := b.db.QueryRow("SELECT id, long_url, short_url FROM URLs WHERE short_url  = $1", key).Scan(
		&u.Id,
		&temp,
		&u.ShortUrl); err != nil {
		return u, false
	}

	o, _ := url.Parse(temp)
	u.LongUrl = *o

	return u, true

}
