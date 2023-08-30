package controllers

import (
	"github.com/Mkkysh/AvitoTest/models"
	"github.com/gofiber/fiber/v2"
)

type SegmentServiceInt interface {
	Add(segment models.Segment, partAuto string) error
	Delete(name string) error
}

type SegmentController struct {
	SegmentService SegmentServiceInt
}

func NewSegmentController(s SegmentServiceInt) *SegmentController {
	return &SegmentController{
		SegmentService: s,
	}
}

func (s *SegmentController) Add(ctx *fiber.Ctx) error {

	var segment models.Segment
	err := ctx.BodyParser(&segment)
	if err != nil {
		return err
	}

	partAuto := ctx.Query("partAuto")

	err = s.SegmentService.Add(segment, partAuto)
	if err != nil {
		return err
	}

	ctx.Status(fiber.StatusCreated)
	return nil
}

func (s *SegmentController) Delete(ctx *fiber.Ctx) error {

	name := ctx.Query("name")
	err := s.SegmentService.Delete(name)
	if err != nil {
		return err
	}

	ctx.Status(fiber.StatusOK)
	return nil
}
