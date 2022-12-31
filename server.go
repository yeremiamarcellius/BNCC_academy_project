package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// func main() {
//  e := echo.New()

//  e.GET("/", func(c echo.Context) error {
//   return c.String(http.StatusOK, "Hello, World!")
//  })
//  e.Logger.Fatal(e.Start(":1323"))

//  e.GET("/users/:id", getUser)
//  // e.POST("/users", saveUser)
//  // e.PUT("/users/:id", updateUser)
//  // e.DELETE("/users/:id", deleteUser)
// }

// func getUser(c echo.Context) error {
//  // User ID from path users/:id
//  id := c.Param("id")
//  return c.String(http.StatusOK, id)
// }

func getUser(c echo.Context) error {
	// User ID from path users/:id
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func welcoming(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func main() {
	e := echo.New()

	e.GET("/", welcoming)
	e.GET("/users/:id", getUser)

	// e.POST("/users", saveUser)
	// e.PUT("/users/:id", updateUser)
	// e.DELETE("/users/:id", deleteUser)
	e.Logger.Fatal(e.Start(":1323"))
}
