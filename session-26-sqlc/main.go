package main

import (
	"flag"
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"log"
	"net/http"

	"session-23-gin-jwt/internal/handlers"
	"session-23-gin-jwt/internal/middlewares"
	"session-23-gin-jwt/internal/repository"
	"session-23-gin-jwt/internal/services"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

func mongoConnect(uri string) (*mongo.Client, error) {
	client, err := mongo.Connect(options.Client().
		ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	return client, nil
}
func sqlxConnect(dsn string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Println("Error connecting to the db", err)
		return nil, err
	}
	return db, nil
}

func sqlORMConnect(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Println(err)
		return nil, err
	}
	return db, nil
}
func main() {
	var dbtype string
	flag.StringVar(&dbtype, "dbtype", "mongodb", "type of database to connect to (e.g. mongodb, postgres)")
	flag.Parse()
	// We will load the env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Cannot find env file")
		return
	}

	r := gin.Default()

	r.Use(cors.Default())

	var repo repository.DbRepository
	fmt.Println(dbtype)
	if dbtype == "mongodb" {
		client, err := mongoConnect("mongodb://localhost:27017/users")
		if err != nil {
			log.Fatal("Error connecting mongodb", err)
			return
		}
		repo = repository.NewMongoRepo(client)
	} else if dbtype == "mysqlx" {
		// import _ "github.com/go-mysql-org/go-mysql/driver"
		// "root:root@localhost:3306?data"
		conn, err := sqlxConnect("root:root@tcp(127.0.0.1:3306)/data?charset=utf8mb4&parseTime=True&loc=Local")
		if err != nil {
			log.Fatal("Error connecting mongodb", err)
			return
		}
		fmt.Println("Initialized repo in sqlx")
		repo = repository.NewMysqlReqo(conn)
	} else if dbtype == "sqlORM" {
		conn, err := sqlORMConnect("root:root@tcp(127.0.0.1:3306)/data?charset=utf8mb4&parseTime=True&loc=Local")
		if err != nil {
			log.Fatal("Error connecting mongodb", err)
			return
		}
		fmt.Println("Initialized repo in sqlORM")
		repo = repository.NewMysqlOrm(conn)
	}

	jwtService := &services.JWTService{}
	handler := handlers.NewHandler(repo, jwtService)
	v1 := r.Group("/api/v1")
	v1.GET("/healthz", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "All good",
		})
	})
	// Two types of groups // auth routes
	auth := v1.Group("/auth") // /api/v1/auth

	auth.POST("/signup", handler.Signup)
	auth.POST("/login", handler.Login)

	// user routes
	user := v1.Group("/user") // /api/v1/user
	user.GET("/getUsers", middlewares.AuthorizationMiddleware(), handler.GetAllUsers)
	user.PUT("/update/:id", middlewares.AuthorizationMiddleware(), handler.UpdateUsers)
	user.DELETE("/delete/:id", middlewares.AuthorizationMiddleware(), handler.DeleteUser)

	err = r.Run("localhost:8080")
	if err != nil {
		return
	}

}
