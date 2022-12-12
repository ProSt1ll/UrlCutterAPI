package main

import (
	"flag"
	"github.com/ProSt1ll/UrlCutterAPI/internal/app/apiserver"
	"github.com/ProSt1ll/UrlCutterAPI/internal/app/saver"
	"log"
	"net/http"
)

var (
	SaveMethod string
	BDPort     string
	ServerPort string
	BDHost     string
	BDName     string
)

func init() {
	flag.StringVar(&SaveMethod, "SaveMethod", "Postgres", "Postgres or inMemory")
	flag.StringVar(&BDPort, "BDPort", "5436", "port of BaseData")
	flag.StringVar(&BDHost, "BDHost", "localhost", "host of BaseData")
	flag.StringVar(&ServerPort, "ServerPort", "8000", "port of server")
	flag.StringVar(&BDName, "BDName", "myurldb", "name of BaseBata")

}
func main() {
	flag.Parse()
	var save saver.Saver
	if SaveMethod == "Postgres" {

		save = saver.NewDB()
		save.Config(BDHost, BDPort, BDName)
		if err := save.Open(); err != nil {
			panic(err)
		}
		defer save.Close()
		if err := save.Open(); err != nil {
			panic(err)
		}
	} else if SaveMethod == "inMemory" {
		save = saver.NewMemSaver()
	} else {
		panic("incorrect save method")
	}

	server := apiserver.New(save)

	err := server.Start()
	if err != nil {
		panic(err)
	}
	log.Println("Server started!")
	log.Fatal(http.ListenAndServe(":"+ServerPort, server.Server))
}
