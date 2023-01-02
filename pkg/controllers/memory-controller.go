package controllers

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo"
	"github.com/youhane/bncc_academy_final/pkg/models"
	"github.com/youhane/bncc_academy_final/pkg/utils"
)

type MemoryController struct{}

func (mc *MemoryController) GetMemories(c echo.Context) error {
	memories := models.GetAllMemories()
	return c.JSON(http.StatusOK, memories)
}

func (mc *MemoryController) CreateMemory(c echo.Context) error {
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

func (mc *MemoryController) GetMemory(c echo.Context) error {
	id := c.Param("id")
	ID, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	memoryDetails, _ := models.GetMemoryById(ID)
	return c.JSON(http.StatusOK, memoryDetails)
}

func (mc *MemoryController) UpdateMemory(c echo.Context) error {
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

func (mc *MemoryController) DeleteMemory(c echo.Context) error {
	id := c.Param("id")
	ID, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	models.DeleteMemory(ID)
	return c.JSON(http.StatusOK, "Memory Deleted")
}
