package main

import (
	"bncc/models"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"bncc/utils"

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

// func getUser(c echo.Context) error {
// 	// User ID from path users/:id
// 	id := c.Param("id")
// 	return c.String(http.StatusOK, id)
// }

// func welcoming(c echo.Context) error {
// 	return c.String(http.StatusOK, "Hello, World!")
// }

func main() {
	e := echo.New()

	e.GET("/", GetMemories)
	e.POST("/memory", CreateMemory)
	e.GET("/memory/:id", GetMemory)
	e.PUT("/memory/:id", UpdateMemory)
	e.DELETE("/memory/:id", DeleteMemory)

	e.Logger.Fatal(e.Start(":9001"))
}

func GetMemories(c echo.Context) error {
	memories := models.GetAllMemories()
	return c.JSON(http.StatusOK, memories)
}

func CreateMemory(c echo.Context) error {
	description := c.FormValue("description")
	title := c.FormValue("title")
	date, _ := time.Parse("2006-01-02", c.FormValue("date"))
	tags := c.FormValue("tags")
	image, err := c.FormFile("image")

	if err != nil {
		fmt.Println("71")
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, err)
	}

	imageName, err := utils.HandleImage(image)
	if err != nil {
		fmt.Println("77")
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	arrayOfTags := strings.Split(tags, ",")
	parsedTags := utils.ParseTags(arrayOfTags)

	memory := models.Memory{
		Title: title,
		Image: imageName,
		Date:  date,
		Desc:  description,
		Tags:  parsedTags,
	}

	memory.CreateMemory()

	return c.JSON(http.StatusOK, memory)
}

func GetMemory(c echo.Context) error {
	id := c.Param("id")
	ID, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	memoryDetails, _ := models.GetMemoryById(ID)
	return c.JSON(http.StatusOK, memoryDetails)
}

func UpdateMemory(c echo.Context) error {
	id := c.Param("id")
	ID, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	memory, _ := models.GetMemoryById(ID)
	memory.Title = c.FormValue("title")
	memory.Desc = c.FormValue("description")
	memory.Date, _ = time.Parse("2006-01-02", c.FormValue("date"))
	tags := c.FormValue("tags")
	arrayOfTags := strings.Split(tags, ",")
	parsedTags := utils.ParseTags(arrayOfTags)
	memory.Tags = parsedTags

	models.UpdateMemory(memory)

	return c.JSON(http.StatusOK, memory)
}

func DeleteMemory(c echo.Context) error {
	id := c.Param("id")
	ID, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	models.DeleteMemory(ID)
	return c.JSON(http.StatusOK, "Memory Deleted")
}
