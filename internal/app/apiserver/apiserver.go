package apiserver

import (
	_ "fmt"
	"github.com/ProSt1ll/UrlCutterAPI/internal/app/saver"
	"github.com/ProSt1ll/UrlCutterAPI/internal/app/urlcut"
	"github.com/ProSt1ll/UrlCutterAPI/model"
	"math"
	"net/http"
	_ "net/http"
	"net/url"
)

type Methods map[string]http.HandlerFunc

type APIServer struct {
	Server     *http.ServeMux
	saver      saver.Saver
	counterUrl int
	cutter     *urlcut.UrlCut
}

func New(saver saver.Saver) *APIServer {
	return &APIServer{
		Server:     http.NewServeMux(),
		saver:      saver,
		counterUrl: 0,
		cutter:     urlcut.New(),
	}
}

func (s *APIServer) Start() error {
	s.Server.Handle("/", RouteMethods(Methods{
		http.MethodGet:  s.ParseGetRequest,
		http.MethodPost: s.ParsePostRequest,
	}))
	return nil
}

func RouteMethods(methods Methods) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resolver, ok := methods[r.Method]
		if !ok {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		resolver(w, r)
	})
}

func (s *APIServer) ParsePostRequest(w http.ResponseWriter, r *http.Request) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	longUrl, err := url.ParseRequestURI(string(body))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if longUrl.Scheme < "ht" {
		longUrl, err = url.Parse(longUrl.Host + longUrl.Path)
		if err != nil {

		}

	} else {
		longUrl, err = url.Parse(longUrl.Scheme + "://" + longUrl.Host + longUrl.Path)
		if err != nil {

		}

	}
	Url, ok := s.saver.LoadShort(*longUrl)
	if ok {
		w.WriteHeader(http.StatusOK)
		_, err = w.Write([]byte("https://ozon.cc/" + Url.ShortUrl + "\n"))
		if err != nil {

		}
		return
	}
	if s.counterUrl == math.MaxInt {
		s.counterUrl = 0
	}
	s.counterUrl++
	Url = model.URLs{}
	Url.ShortUrl = s.cutter.CreateShortURL(s.counterUrl)
	Url.LongUrl = *longUrl
	s.saver.StoreURL(Url)
	_, err = w.Write([]byte("https://ozon.cc/" + Url.ShortUrl + "\n"))
	if err != nil {

	}
	w.WriteHeader(http.StatusOK)
	return
}

func (s *APIServer) ParseGetRequest(w http.ResponseWriter, r *http.Request) {
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	shortUrl, err := url.Parse(string(body))
	if err != nil {

	}
	url, ok := s.saver.LoadLong(shortUrl.Path[1:])
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(url.LongUrl.String() + "\n"))
}
