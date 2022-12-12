package model

import (
	"net/url"
)

type URLs struct {
	LongUrl  url.URL
	ShortUrl string
	Id       int
}

type Config struct {
	UrlCnt int
	List   string
}
