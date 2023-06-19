package main

import (
	"flag"
	"os"

	"github.com/ifty123/simple_online_store/database"
	"github.com/ifty123/simple_online_store/database/migration"
	"github.com/ifty123/simple_online_store/database/seeder"
	"github.com/ifty123/simple_online_store/internal/factory"
	"github.com/ifty123/simple_online_store/internal/http"
	"github.com/ifty123/simple_online_store/internal/middleware"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	database.GetConnection()
}

func main() {
	database.CreateConnection()

	var m string // for check migration
	var s string // for check seeder

	flag.StringVar(
		&m,
		"m",
		"none",
		`this argument for check if user want to migrate table, rollback table, or status migration
to use this flag:
	use -m=migrate for migrate table
	use -m=rollback for rollback table
	use -m=status for get status migration`,
	)

	flag.StringVar(
		&s,
		"s",
		"none",
		`this argument for check if user want to seed table
to use this flag:
	use -s=all to seed all table`,
	)

	flag.Parse()

	if m == "migrate" {
		migration.Migrate()
	} else if m == "rollback" {
		migration.Rollback()
	} else if m == "status" {
		migration.Status()
	}

	if s == "all" {
		seeder.NewSeeder().DeleteAll()
		seeder.NewSeeder().SeedAll()
	}

	//factory database
	f := factory.NewFactory()
	e := echo.New()

	e.Use(middleware.CORS)
	middleware.LogMiddlewares(e)

	http.NewHttp(e, f)

	e.Logger.Fatal(e.Start(":" + os.Getenv("APP_PORT")))
}
