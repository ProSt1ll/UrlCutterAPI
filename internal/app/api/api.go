package api

import (
	"github.com/ProSt1ll/UrlCutterAPI/internal/app/saver"
	"github.com/ProSt1ll/UrlCutterAPI/internal/app/urlcut"
	"github.com/ProSt1ll/UrlCutterAPI/model"
	"math"
	"net/http"
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
		Server: http.NewServeMux(),
		saver:  saver,
		cutter: urlcut.New(),
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
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	longUrl, err := url.ParseRequestURI(string(body))
	if err != nil {
		//if URL invalid, example: daskgjflajsfpdgjld.ru
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if longUrl.Scheme > "" {
		longUrl, err = url.Parse(longUrl.Scheme + "://" + longUrl.Host + longUrl.Path)
		if err != nil {
			panic("parse error")
		}
	} else {
		longUrl, err = url.Parse(longUrl.Host + longUrl.Path)
		if err != nil {
			panic("parse error")
		}

	}
	// try to find in database
	Url, ok := s.saver.LoadShort(*longUrl)
	if ok {
		w.WriteHeader(http.StatusOK)
		_, err = w.Write([]byte("https://ozon.cc/" + Url.ShortUrl + "\n"))
		if err != nil {
			panic(err)
		}
		return
	}

	// if URLs more than can be in variable
	if s.counterUrl == math.MaxInt {
		s.counterUrl = 0
	}
	s.counterUrl++

	Url = model.URLs{
		ShortUrl: s.cutter.CreateShortURL(s.counterUrl),
		LongUrl:  *longUrl,
	}
	var n int
	if n, err = s.saver.StoreURL(Url); err != nil {
		panic(err)
	}
	//if we save config of urlcutter, we can continue
	if n > 0 {
		s.counterUrl = n
	}
	//create URL
	_, err = w.Write([]byte("https://ozon.cc/" + Url.ShortUrl + "\n"))
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
}

func (s *APIServer) ParseGetRequest(w http.ResponseWriter, r *http.Request) {
	body := make([]byte, r.ContentLength)
	//body.read always give a EOF error
	if n, err := r.Body.Read(body); err != nil && n < 1 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	shortUrl, err := url.Parse(string(body))
	if err != nil {
		panic("parse error")
	}

	loadedURL, ok := s.saver.LoadLong(shortUrl.Path[1:])

	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte(loadedURL.LongUrl.String() + "\n")); err != nil {
		panic(err)
	}
}
