package main

import (
	"database/sql"
	"fmt"
	"go_api_sample/infrastructure"
	"go_api_sample/interfaces"
	"go_api_sample/usecase"

	"log"
	"os"

	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

func main() {
	db := dbConnect()

	defer db.Close()

	e := echo.New()

	userRepository := infrastructure.NewUserRepositoryInfrastructure(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := interfaces.NewUserController(userUsecase)
	controllers := interfaces.NewControllers(userController)
	controllers.Mount(e.Group(""))

	e.Logger.Fatal(e.Start(":1323"))
}

func dbConnect() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbconf := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE_NAME"),
	)

	db, err := sql.Open("mysql", dbconf)
	if err != nil {
		fmt.Println(err.Error())
	}

	return db
}
