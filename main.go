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
	"tani-hub-v2/controller"
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

	//article
	router.GET("/api/article", controller.GetAllArticle)
	router.GET("/api/article/:id", controller.GetArticleById)
	router.POST("/api/article", controller.InsertArticle)
	router.PUT("/api/article/:id", controller.UpdateArticle)
	router.DELETE("/api/article/:id", controller.DeleteArticle)

	//category
	router.GET("/api/category", controller.GetAllCategory)
	router.GET("/api/category/:id", controller.GetCategoryById)
	router.POST("/api/category", controller.InsertCategory)
	router.PUT("/api/category/:id", controller.UpdateCategory)
	router.DELETE("/api/category/:id", controller.DeleteCategory)

	//category
	router.GET("/api/product", controller.GetAllProduct)
	router.GET("/api/product/:id", controller.GetProductById)
	router.POST("/api/product", controller.InsertProduct)
	router.PUT("/api/product/:id", controller.UpdateProduct)
	router.DELETE("/api/product/:id", controller.DeleteProduct)

	//order
	router.POST("/api/order", controller.InsertOrder)
	router.GET("/api/order", controller.GetAllOrder)

	errRun := router.Run(":" + portApp)
	if errRun != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
