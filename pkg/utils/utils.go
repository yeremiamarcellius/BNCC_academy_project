package utils

import (
	"io"
	"mime/multipart"
	"os"

	"github.com/youhane/bncc_academy_final/pkg/models"
)

func HandleImage(img *multipart.FileHeader) (string, error) {
	src, err := img.Open()
	if err != nil {
		return "Error", err
	}
	defer src.Close()

	imgPath := "public/images/" + img.Filename
	dest, err := os.Create(imgPath)
	if err != nil {
		return "Error", err
	}
	defer dest.Close()

	_, err = io.Copy(dest, src)
	if err != nil {
		return "Error", err
	}

	return imgPath, nil
}

func ParseTags(tags []string) []models.Tag {
	var parsedTags []models.Tag
	for _, tag := range tags {
		tag := models.Tag{Name: tag}
		tag.CreateTag()
		parsedTags = append(parsedTags, tag)
	}
	return parsedTags
}
