package controllers

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/Mkkysh/AvitoTest/dto"
	"github.com/Mkkysh/AvitoTest/models"
	"github.com/gofiber/fiber/v2"
)

type LogServiceInt interface {
	Add(logs []models.Log) error
	Get(date string) ([]dto.LogResponse, error)
}

type LogController struct {
	LogService LogServiceInt
}

func NewLogController(s LogServiceInt) *LogController {
	return &LogController{
		LogService: s,
	}
}

func (c *LogController) Get(ctx *fiber.Ctx) error {

	date := ctx.Query("date")

	logs, err := c.LogService.Get(date)

	if err != nil {
		return err
	}

	saveDir := "tmpCsv"

	filePath := fmt.Sprintf("%s/log_%s.csv", saveDir, date)

	file, err := os.Create(filePath)

	if err != nil {
		return err
	}

	writer := csv.NewWriter(file)

	header := []string{"IdUser", "NameSegment", "Operation", "Timestamp"}
	err = writer.Write(header)
	if err != nil {
		return err
	}

	for _, log := range logs {
		data := []string{
			strconv.Itoa(log.IdUser),
			log.NameSegment,
			log.Operation,
			log.Timestamp,
		}

		err = writer.Write(data)
		if err != nil {
			return err
		}
	}

	if err := writer.Error(); err != nil {
		return err
	}

	writer.Flush()

	ctx.Type("text/csv")
	ctx.Set("Content-Disposition", "attachment; filename=log.csv")

	ctx.Status(fiber.StatusOK)
	ctx.Download(filePath)

	defer func() {

		err = file.Close()
		if err != nil {
			log.Printf("Failed to remove file: %s", err)
		}
		err = os.Remove(filePath)
		if err != nil {
			log.Printf("Failed to remove file: %s", err)
		}
	}()

	return nil
}
