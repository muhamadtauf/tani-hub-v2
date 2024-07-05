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
	"tani-hub-v2/middleware"
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

	//auth
	router.POST("/api/signup", controller.Signup)
	router.POST("/api/signup/admin", middleware.RequireAuthAdmin, controller.SignupAdmin)
	router.POST("/api/login", controller.Login)

	//article
	router.GET("/api/article", controller.GetAllArticle)
	router.GET("/api/article/:id", controller.GetArticleById)
	router.POST("/api/article", middleware.RequireAuthAdmin, controller.InsertArticle)
	router.PUT("/api/article/:id", middleware.RequireAuthAdmin, controller.UpdateArticle)
	router.DELETE("/api/article/:id", middleware.RequireAuthAdmin, controller.DeleteArticle)

	//category
	router.GET("/api/category", controller.GetAllCategory)
	router.GET("/api/category/:id", middleware.RequireAuthAdmin, controller.GetCategoryById)
	router.POST("/api/category", middleware.RequireAuthAdmin, controller.InsertCategory)
	router.PUT("/api/category/:id", middleware.RequireAuthAdmin, middleware.RequireAuthAdmin, controller.UpdateCategory)
	router.DELETE("/api/category/:id", middleware.RequireAuthAdmin, middleware.RequireAuthAdmin, controller.DeleteCategory)

	//product
	router.GET("/api/product", controller.GetAllProduct)
	router.GET("/api/product/:id", controller.GetProductById)
	router.POST("/api/product", middleware.RequireAuthAdmin, controller.InsertProduct)
	router.PUT("/api/product/:id", middleware.RequireAuthAdmin, controller.UpdateProduct)
	router.DELETE("/api/product/:id", middleware.RequireAuthAdmin, controller.DeleteProduct)

	//order
	router.POST("/api/order", middleware.RequireAuthUser, controller.InsertOrder)
	router.GET("/api/order", middleware.RequireAuthAdmin, controller.GetAllOrder)
	router.GET("/api/order/:code", middleware.RequireAuth, controller.GetOrderByCode)
	router.GET("/api/order/user/:user_id", middleware.RequireAuth, controller.GetOrderByUserId)
	router.PUT("/api/order/processed/:id", middleware.RequireAuthAdmin, controller.UpdateOrderToProcessed)
	router.PUT("/api/order/shipped/:id", middleware.RequireAuthAdmin, controller.UpdateOrderToShipped)
	router.PUT("/api/order/finished/:id", middleware.RequireAuthAdmin, controller.UpdateOrderToFinished)

	errRun := router.Run(":" + portApp)
	if errRun != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
