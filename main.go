package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
	"strconv"
	"tani-hub-v2/database"
)

var DB *sql.DB

func main() {
	err := godotenv.Load("config/.env")
	if err != nil {
		fmt.Println("failed load file environment")
	} else {
		fmt.Println("success read file environment")
	}

	port, _ := strconv.Atoi(os.Getenv("PGPORT"))
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("PGHOST"), port, os.Getenv("PGUSER"), os.Getenv("PGPASSWORD"), os.Getenv("PGDATABASE"))

	DB, err = sql.Open("postgres", psqlInfo)
	err = DB.Ping()
	if err != nil {
		fmt.Println("DB Connection Failed")
		panic(err)
	} else {
		fmt.Println("DB Connection Success")
	}

	database.DbMigrate(DB)

	defer DB.Close()

	portApp := os.Getenv("PORT")
	if portApp == "" {
		log.Fatal("PORT environment variable not set")
	}

	//router
	router := gin.Default()

	errRun := router.Run(":" + portApp)
	if errRun != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
