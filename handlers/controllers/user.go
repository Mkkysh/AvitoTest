package controllers

import (
	"strconv"

	"github.com/Mkkysh/AvitoTest/dto"
	"github.com/gofiber/fiber/v2"
)

type UserServiceInt interface {
	UpdateSegment(id int, AddSegments []interface{}, RemoveSegments []interface{}) error

	GetSegemnts(id int) ([]interface{}, error)
}

type UserController struct {
	UserService UserServiceInt
}

func NewUserController(u UserServiceInt) *UserController {
	return &UserController{
		UserService: u,
	}
}

func (u *UserController) UpdateSegment(ctx *fiber.Ctx) error {

	idStr := ctx.Params("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return err
	}

	var dto dto.ChangeSegments
	err = ctx.BodyParser(&dto)
	if err != nil {
		return err
	}

	//log.Println(dto.AddSegments)

	err = u.UserService.UpdateSegment(id, dto.AddSegments, dto.RemoveSegments)
	if err != nil {
		return err
	}

	ctx.Status(fiber.StatusOK)

	return nil
}

func (c *UserController) GetSegemnts(ctx *fiber.Ctx) error {
	idStr := ctx.Params("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return err
	}

	segments, err := c.UserService.GetSegemnts(id)
	if err != nil {
		return err
	}

	if len(segments) == 0 {
		ctx.Status(fiber.StatusNotFound)
		return nil
	}

	err = ctx.JSON(segments)
	if err != nil {
		return err
	}

	ctx.Status(fiber.StatusOK)

	return nil
}
