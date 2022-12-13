package main

import (
	"flag"
	"github.com/ProSt1ll/UrlCutterAPI/internal/app/api"
	"github.com/ProSt1ll/UrlCutterAPI/internal/app/saver"
	"log"
	"net/http"
	"os"
)

var (
	SaveMethod string
	DBPort     string
	ServerPort string
	DBHost     string
	DBName     string
)

func init() {
	//input without docker
	flag.StringVar(&SaveMethod, "SaveMethod", "inMemory", "postgres or inMemory")
	flag.StringVar(&DBPort, "DBPort", "5432", "port of BaseData")
	flag.StringVar(&DBHost, "DBHost", "localhost", "host of BaseData")
	flag.StringVar(&ServerPort, "ServerPort", "8000", "port of server")
	flag.StringVar(&DBName, "DBName", "myurldb", "name of BaseBata")
	flag.Parse()
}

func main() {
	//input for docker
	Docker := os.Getenv("Docker")
	if Docker == "true" {
		SaveMethod = os.Getenv("SaveMethod")
		DBPort = os.Getenv("DBPort")
		DBHost = os.Getenv("DBHost")
		ServerPort = os.Getenv("ServerPort")
		DBName = os.Getenv("DBName")
	}

	var save saver.Saver
	log.Println(SaveMethod)
	switch SaveMethod {
	case "postgres":
		db := saver.NewDB(DBHost, DBPort, DBName)
		if err := db.Open(); err != nil {
			log.Fatal("incorrect save method")
		}
		defer func() {
			if err := db.Close(); err != nil {
				log.Fatal(err)
			}
		}()
		save = saver.NewDBSaver(db.DB)
	case "inMemory":
		save = saver.NewMemSaver()
	default:
		log.Fatal("incorrect save method")
	}

	server := api.New(save)
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
	log.Println("Starting server...")
	log.Println("localhost" + ":" + ServerPort)
	log.Fatal(http.ListenAndServe(":"+ServerPort, server.Server))

}
