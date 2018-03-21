package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Data
var Routes = map[string]string{
	"a": "ra",
	"b": "rb",
}

// Helpers
func request(url string) string {

	response, err := http.Get(url)

	if err != nil {

		fmt.Printf("%s", err)

	} else {

		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)

		if err != nil {
			fmt.Printf("%s", err)
		}

		return string(contents)
	}

	return ""
}

// Routes
func setupRouter() *gin.Engine {

	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	// Dynamic routing.
	r.GET("/dynamic/:route", func(c *gin.Context) {
		route := c.Params.ByName("route")
		value, ok := Routes[route]

		if ok {

			c.JSON(200, gin.H{"route": route, "value": value})

		} else {

			c.JSON(200, gin.H{"route": route, "status": "no value"})
		}
	})

	r.GET("/recurse/:route", func(c *gin.Context) {
		route := c.Params.ByName("route")

		i := len(route)
		prefix := route[0:1]
		var res string

		if i > 1 {

			res = request("http://localhost:8080/recurse/" + route[1:])

		} else {

			res = "|"
		}

		c.String(200, prefix + "-" + res)
	})

	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080") // Listen and Server in 0.0.0.0:8080
}
