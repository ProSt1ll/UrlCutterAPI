package saver

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ProSt1ll/UrlCutterAPI/model"
	"github.com/stretchr/testify/assert"
	"log"
	"net/url"
	"regexp"
	"testing"
)

func TestDBSaver_StoreURL(t *testing.T) {
	db, mock, err := sqlmock.New()
	r := NewDBSaver(db)
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		return
	}

	type args struct {
		item model.URLs
	}

	testURL, err := url.Parse("exampleLong")
	if err != nil {

	}

	testTable := []struct {
		name    string
		input   args
		mock    func()
		want    int
		wantErr bool
	}{
		{
			name: "Ok",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
				mock.ExpectQuery(regexp.QuoteMeta("INSERT INTO urls (long_url,short_url) VALUES ($1,$2) RETURNING id")).
					WithArgs("exampleLong", "exampleShort").WillReturnRows(rows).WillReturnError(nil)

			},
			input: args{
				item: model.URLs{
					LongUrl:  *testURL,
					ShortUrl: "exampleShort",
				},
			},
			want: 1,
		},
		{
			name: "Error",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
				mock.ExpectQuery(regexp.QuoteMeta("INTO urls (long_url,short_url) VALUES ($1,$2) RETURNING id")).
					WithArgs("exampleLong", "exampleShort").WillReturnRows(rows).WillReturnError(err)

			},
			input: args{
				item: model.URLs{
					LongUrl:  *testURL,
					ShortUrl: "exampleShort",
				},
			},
			want: 1,
		},
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			got, err := r.StoreURL(tt.input.item)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestDBSaver_LoadLong(t *testing.T) {
	db, mock, err := sqlmock.New()
	r := NewDBSaver(db)
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		return
	}

	type args struct {
		item string
	}

	testURL, err := url.Parse("exampleLong")
	if err != nil {

	}

	testTable := []struct {
		name    string
		input   args
		mock    func()
		want    model.URLs
		wantErr bool
		ok      bool
	}{
		{
			name: "Ok",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "long_url", "short_url"})
				for i := 0; i < 10; i++ {
					rows.AddRow("1", "exampleLong", "exampleShort")
				}
				mock.ExpectQuery(regexp.QuoteMeta("SELECT id, long_url, short_url FROM URLs WHERE short_url  = $1")).
					WithArgs("exampleShort").WillReturnRows(rows).WillReturnError(nil)

			},
			input: args{
				item: "exampleShort",
			},
			want: model.URLs{
				LongUrl:  *testURL,
				ShortUrl: "exampleShort",
				Id:       1,
			},
			ok: true,
		},
		{
			name: "Error",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
				mock.ExpectQuery(regexp.QuoteMeta("INTO urls (long_url,short_url) VALUES ($1,$2) RETURNING id")).
					WithArgs("exampleLong", "exampleShort").WillReturnRows(rows).WillReturnError(err)

			},
			input: args{
				item: "example2Short",
			},
			want: model.URLs{},
			ok:   false,
		},
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			got, ok := r.LoadLong(tt.input.item)

			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, ok, tt.ok)
		})
	}
}

func TestDBSaver_LoadShort(t *testing.T) {
	db, mock, err := sqlmock.New()
	r := NewDBSaver(db)
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		return
	}

	type args struct {
		item url.URL
	}

	testURL, err := url.Parse("exampleLong")
	if err != nil {

	}

	test2URL, err := url.Parse("example2Long")
	if err != nil {

	}

	testTable := []struct {
		name    string
		input   args
		mock    func()
		want    model.URLs
		wantErr bool
		ok      bool
	}{
		{
			name: "Ok",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "long_url", "short_url"})
				for i := 0; i < 10; i++ {
					rows.AddRow("1", "exampleLong", "exampleShort")
				}
				mock.ExpectQuery(regexp.QuoteMeta("SELECT id, long_url, short_url FROM urls WHERE long_url  = $1")).
					WithArgs("exampleLong").WillReturnRows(rows)

			},
			input: args{
				item: *testURL,
			},
			want: model.URLs{
				LongUrl:  *testURL,
				ShortUrl: "exampleShort",
				Id:       1,
			},
			ok: true,
		},
		{
			name: "Error",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "long_url", "short_url"})
				for i := 0; i < 10; i++ {
					rows.AddRow("1", "exampleLong", "exampleShort")
				}
				mock.ExpectQuery(regexp.QuoteMeta("SELECT id, long_url, short_url FROM urls WHERE long_url  = $1")).
					WithArgs("exampleLong").WillReturnRows(rows)

			},
			input: args{
				item: *test2URL,
			},
			want: model.URLs{},
			ok:   false,
		},
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			got, ok := r.LoadShort(tt.input.item)

			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, ok, tt.ok)
		})
	}
}
