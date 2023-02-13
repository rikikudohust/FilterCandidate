package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	// "sync"

	_ "io/ioutil"
	_ "net/http"

	"github.com/rikikudohust/FilterCandidate/api"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"
)

var (
	port                                                           int
	host, userRead, userWrite, passwordRead, passwordWrite, dbName string
)

func loadConfig() {
	//load host
	host = os.Getenv("HOST")
	if host == "" {
		host = "localhost"
	}
	// load port
	port, _ = strconv.Atoi(os.Getenv("PORT"))
	if port == 0 {
		port = 5432
	}
	// read user
	userRead = os.Getenv("UserRead")
	if userRead == "" {
		log.Fatal("Invalid read user")
	}

	passwordRead = os.Getenv("PasswordRead")

	// write user
	userWrite = os.Getenv("UserWrite")
	if userWrite == "" {
		log.Fatal("Invalid write user")
	}
	passwordWrite = os.Getenv("PasswordWrite")
	// load db
	dbName = os.Getenv("DBName")
}

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	loadConfig()

	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
	}

	// fmt.Println(data)
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "accept", "origin", "Cache-Control", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 15 * time.Second,
	}))

	setup := api.Config{
		Version: "Test",
		Server:  router,
	}
	_, err = api.NewAPI(setup)
	if err != nil {
		fmt.Println("error")
	}

	router.Run(":5000")
}
