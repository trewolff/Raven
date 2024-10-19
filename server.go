package main

import (
	"flag"
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {
	echoServer := echo.New()
	var port = flag.Int("port", 8080, "Port for test HTTP server")
	echoServer.Start(fmt.Sprintf("0.0.0.0:%d", *port))
}
