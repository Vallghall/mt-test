package main

import (
	"github.com/Vallghall/mt/test/internal/handler"
	"github.com/Vallghall/mt/test/internal/service"
	"github.com/Vallghall/mt/test/internal/storage"
	"github.com/Vallghall/mt/test/internal/storage/neo"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln(err)
	}
	session, closeDriver := neo.NewNeo4jSession(&neo.Config{
		URI:      os.Getenv("NEO4J_URI"),
		Username: os.Getenv("NEO4J_USER"),
		Password: os.Getenv("NEO4J_PW"),
	})
	defer closeDriver()
	defer session.Close()

	stor := storage.NewStorage(session)
	stor.LoadGraphData(os.Getenv("SRC_FILE"))
	serv := service.NewService(stor)
	h := handler.NewHandler(serv)
	log.Fatalln(h.InitRoutes().Run())
}
