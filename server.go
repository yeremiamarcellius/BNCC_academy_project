package main

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/yeremiamarcellius/bncc_academy_project/models"
	"github.com/yeremiamarcellius/bncc_academy_project/utils"
)

func main() {
	e := echo.New()

	// c := &controllers.MemoryController{}

	e.Use(middleware.CORS())
	e.Static("/", "public")
	e.GET("/", GetMemories)
	e.POST("/memory", CreateMemory)
	e.GET("/memory/:id", GetMemory)
	e.PUT("/memory/:id", UpdateMemory)
	e.DELETE("/memory/:id", DeleteMemory)

	e.POST("/login", Login)
	e.POST("/register", Register)

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
		return c.JSON(http.StatusBadRequest, err)
	}

	imageName, err := utils.HandleImage(image)
	if err != nil {
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

func Login(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	user, err := models.GetUserByEmail(email)

	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	if user.Password != password {
		return c.JSON(http.StatusUnauthorized, "Wrong password")
	}

	token, err := models.GenerateJWT(email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token":    token,
		"username": user.Username,
		"email":    user.Email,
	})
}

func Register(c echo.Context) error {
	username := c.FormValue("username")
	email := c.FormValue("email")
	password := c.FormValue("password")

	user := models.User{
		Username: username,
		Email:    email,
		Password: password,
	}

	user.CreateUser()

	return c.JSON(http.StatusOK, user)
}

func GetMemoriesByUserId(c echo.Context) error {
	id := c.Param("id")
	ID, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	user, err := models.GetUserById(ID)
	if err != nil {
		return c.JSON(http.StatusNotFound, "User not found")
	}

	memories := models.GetAllMemoriesByUser(user)
	return c.JSON(http.StatusOK, memories)
}
