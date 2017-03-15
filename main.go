// This small program is just a small web server created in static mode
// in order to provide the smallest docker image possible

package main // import "github.com/firstderm/goStatic"

import (
	"flag"
	"log"
	"strconv"

	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
)

var (
	// Def of flags
	portPtr  = flag.Int("p", 8043, "The listening port")
	pathPtr  = flag.String("static", "/srv/http", "The path for the static files")
	html5Ptr = flag.Bool("html5", false, "If enabled, all non-matching routes will be directed to the index page")
)

func main() {

	flag.Parse()

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(mw.Logger())
	e.Use(mw.Recover())

	// Routes
	e.Use(mw.StaticWithConfig(mw.StaticConfig{
		Root:  *pathPtr,
		HTML5: true,
	}))

	log.Println("Starting goStatic")

	port := ":" + strconv.FormatInt(int64(*portPtr), 10)

	// Start server with unsecure HTTP
	log.Println("Starting serving", *pathPtr, "on", *portPtr)
	e.Start(port)
}
