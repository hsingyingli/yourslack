package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	db "github.com/hsingyingli/yourslack-backend/db/sqlc"
	"github.com/hsingyingli/yourslack-backend/pkg/server"
	"github.com/hsingyingli/yourslack-backend/pkg/utils"
	_ "github.com/lib/pq"
)

func main() {
	config, err := utils.LoadConfig()

	if err != nil {
		log.Fatal("Unable to read config file")
	}
	gin.SetMode(config.GIN_MODE)

	dbSource := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		config.DB_USERNAME,
		config.DB_PASSWORD,
		config.DB_URL,
		"5432",
		config.DB_TABLE,
	)

	dbName := "postgres"
	conn, err := sql.Open(dbName, dbSource)

	if err != nil {
		log.Fatal("Unable to connect to db")
	}

	store := db.NewStore(conn)
	server := server.NewServer(config, store)
	log.Fatal(server.Start(config.PORT))
}
