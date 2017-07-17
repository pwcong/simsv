package main

import (
	"log"
	"os"
	"path/filepath"

	"strconv"

	"github.com/BurntSushi/toml"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type config struct {
	IP   string
	Port int
	Root string
}

var Config config

const DEFAULT_CONFIG = `
ip = "0.0.0.0"
port = 6666
root = "www"
`

func init() {

	_, err := toml.DecodeFile(filepath.Join(filepath.Dir(os.Args[0]), "simsv.conf"), &Config)

	if err == nil {

		log.Print("Custom configutation has been loaded successfully.")

	} else {
		_, err := toml.Decode(DEFAULT_CONFIG, &Config)
		if err != nil {
			log.Fatal(err.Error())
		} else {
			log.Print("Failed to load custom configuration. Use default.")
		}
	}
}

func initMiddlewares(e *echo.Echo) {

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{}))

}

func initRoutes(e *echo.Echo) {

	e.Static("/", Config.Root)

}

func main() {

	e := echo.New()
	initMiddlewares(e)
	initRoutes(e)

	e.Logger.Fatal(e.Start(Config.IP + ":" + strconv.Itoa(Config.Port)))

}
