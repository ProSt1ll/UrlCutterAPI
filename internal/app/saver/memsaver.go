package saver

import (
	"github.com/ProSt1ll/UrlCutterAPI/model"
	"net/url"
	"sync"
)

type MemSaver struct {
	cacheShort map[url.URL]string
	cacheLong  map[string]url.URL
	mutex      sync.RWMutex
}

func NewMemSaver() Saver {
	return &MemSaver{
		cacheShort: make(map[url.URL]string),
		cacheLong:  make(map[string]url.URL),
		mutex:      sync.RWMutex{},
	}
}

func (m *MemSaver) StoreURL(urls model.URLs) (int, error) {
	defer m.mutex.Unlock()
	m.mutex.Lock()
	m.cacheLong[urls.ShortUrl] = urls.LongUrl
	m.cacheShort[urls.LongUrl] = urls.ShortUrl

	return 0, nil
}

func (m *MemSaver) LoadShort(key url.URL) (model.URLs, bool) {
	defer m.mutex.RUnlock()
	m.mutex.RLock()

	return model.URLs{
		LongUrl:  key,
		ShortUrl: m.cacheShort[key],
	}, !(m.cacheShort[key] == "")
}

func (m *MemSaver) LoadLong(key string) (model.URLs, bool) {
	defer m.mutex.RUnlock()
	m.mutex.RLock()
	a := m.cacheLong[key]

	return model.URLs{
		LongUrl:  m.cacheLong[key],
		ShortUrl: key,
	}, !(a.String() == "")
}

func (m *MemSaver) Close() error {
	return nil
}
