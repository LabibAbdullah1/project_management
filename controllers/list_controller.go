package controllers

import (
	"ProjectManagement/models"
	"ProjectManagement/services"
	"ProjectManagement/utils"

	"github.com/gofiber/fiber/v2"
)

type ListController struct {
	service services.ListService
}

func NewListController(listService services.ListService) *ListController {
	return &ListController{service: listService}
}

func (c *ListController) CreateList(ctx *fiber.Ctx) error {
	list := new(models.List)
	if err := ctx.BodyParser(list); err != nil {
		return utils.BadRequest(ctx, "Gagal Membaca Request", err.Error())
	}
	err := c.service.Create(list)
	if err != nil {
		return utils.BadRequest(ctx, "Gagal Membuat List", err.Error())
	}

	return utils.Created(ctx, "List berhasil dibuat", list)
}
