package main

import (
	"fmt"
	"gotodo/controller"
	"gotodo/lib"
	"gotodo/lib/util"
	"io/ioutil"
	"log"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	sane "gitlab.com/bloom42/sane/sane-go"
	"gopkg.in/go-playground/validator.v9"
)

func main() {
	fmt.Println("Hi!")

	// init zap logger
	logger := util.InitLogger()
	defer logger.Sync()

	// To open a file
	file, err := os.Open("config.txt") // For read access
	if err != nil {
		log.Printf("Error:%#v", err)
		return
	}

	// To now read the file as a byte slice
	byteSlice, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println(err)
		return
	}

	var dbi util.ConfigFile

	err = sane.Unmarshal(byteSlice, &dbi)
	if err != nil {
		log.Println(err)
		return
	}

	db := lib.InitDB(dbi.DBInfo.User, dbi.DBInfo.Password, dbi.DBInfo.DBName)
	defer db.Close()

	e := echo.New()
	e.File("/", "public/index.html")

	lc := middleware.LoggerConfig{
		Format: `[${method}] ${status} - ${uri}` +
			` - ${latency_human}, rx:${bytes_in}, tx:${bytes_out}` + "\n",
	}

	e.Use(middleware.LoggerWithConfig(lc))

	e.Use(middleware.Recover())

	e.Validator = &util.CustomValidator{
		Validator: validator.New(),
	}

	controller.DefineRoutes(e, db, logger, "/api")
	logger.Debug("routes defined")

	port := fmt.Sprintf(":%s", dbi.Port)
	e.Logger.Fatal(e.Start(port))

}
