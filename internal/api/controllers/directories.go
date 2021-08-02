package controllers

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/xolisamatika/go-rest-api/internal/pkg/models"
)

func GetDirectoryList(c *fiber.Ctx) error {

	directory := "./"
	// directory := c.Params("directory")
	fmt.Println("Input:", directory)
	files, err := ioutil.ReadDir(directory)
	path, _ := filepath.Abs(directory)
	if err != nil {
		log.Fatal(err)
	}
	var filesResponse []models.Directory = make([]models.Directory, 0)
	for _, file := range files {
		fmt.Println(filepath.Join(path, file.Name()))

		data := new(models.Directory)
		data.FullPath = filepath.Join(path, file.Name())
		data.Size = file.Size()
		data.IsDirective = file.IsDir()
		data.ModificationTime = file.ModTime()

		filesResponse = append(filesResponse, *data)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"directoryList": filesResponse,
		},
	})
}
