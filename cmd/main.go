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
	flag.StringVar(&SaveMethod, "SaveMethod", "postgres", "postgres or inMemory")
	flag.StringVar(&DBPort, "DBPort", "5432", "port of BaseData")
	flag.StringVar(&DBHost, "DBHost", "localhost", "host of BaseData")
	flag.StringVar(&ServerPort, "ServerPort", "8000", "port of server")
	flag.StringVar(&DBName, "DBName", "myurldb", "name of BaseBata")

}
func main() {
	flag.Parse()
	var save saver.Saver

	//input for docker
	Docker := os.Getenv("Docker")
	if Docker == "true" {
		SaveMethod = os.Getenv("SaveMethod")
		DBPort = os.Getenv("DBPort")
		DBHost = os.Getenv("DBHost")
		ServerPort = os.Getenv("ServerPort")
		DBName = os.Getenv("DBName")
	}

	switch SaveMethod {
	case "postgres":
		db := saver.NewDB(DBHost, DBPort, DBName)
		if err := db.Open(); err != nil {
			log.Fatal("incorrect save method")
		}
		defer func() {
			if err := db.Close(); err != nil {
				log.Fatalf("close db error", err)
			}
		}()
		save = saver.NewDBSaver(db.DB)
	case "inMemory":
		save = saver.NewMemSaver()
	default:
		log.Fatal("incorrect save method")
	}

	server := api.New(save)
	err := server.Start()
	if err != nil {
		panic(err)
	}
	log.Println("Server started!")
	log.Fatal(http.ListenAndServe(":"+ServerPort, server.Server))
}
