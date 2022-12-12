package saver

import (
	"github.com/ProSt1ll/UrlCutterAPI/model"
	"net/url"
)

type Saver interface {
	StoreURL(model.URLs) (int, error)
	LoadShort(url.URL) (model.URLs, bool)
	LoadLong(string) (model.URLs, bool)
	Open() error
	Close() error
	Config(string, string, string)
}
