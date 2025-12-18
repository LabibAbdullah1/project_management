package controllers

import (
	"ProjectManagement/models"
	"ProjectManagement/services"
	"ProjectManagement/utils"

	"github.com/jinzhu/copier"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	service services.UserService
}

func NewUserController(s services.UserService) *UserController {
	return &UserController{service: s}
}

func (c *UserController) Register(ctx *fiber.Ctx) error {
	user := new(models.User)

	if err := ctx.BodyParser(user); err != nil {
		return utils.BadRequest(ctx, "Gagal parsing data", err.Error())
	}

	if err := c.service.Register(user); err != nil {
		return utils.BadRequest(ctx, "Gagal registrasi user", err.Error())
	}
	var UserResp models.UserResponse
	_ = copier.Copy(&UserResp, &user)
	return utils.Created(ctx, "User berhasil dibuat", UserResp)
}

func (c *UserController) Login(ctx *fiber.Ctx) error {
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := ctx.BodyParser(&body); err != nil {
		return utils.BadRequest(ctx, "Invalid Request", err.Error())
	}
	user, err := c.service.Login(body.Email, body.Password)
	if err != nil {
		return utils.Unauthorized(ctx, "Login Failed", err.Error())
	}
	token, _ := utils.GenerateToken(user.InternalID, user.Role, user.Email, user.PublicID)
	refreshToken, _ := utils.GenerateRefreshToken(user.InternalID, user.Role, user.Email, user.PublicID)

	var UserResp models.UserResponse
	_ = copier.Copy(&UserResp, &user)
	return utils.Success(ctx, "Login Successful", fiber.Map{
		"access_token":  token,
		"refresh_token": refreshToken,
		"user":          UserResp,
	})
}

func (c *UserController) GetUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	user, err := c.service.GetByPublicID(id)
	if err != nil {
		return utils.NotFound(ctx, "User not found", err.Error())
	}
	var UserResp models.UserResponse
	_ = copier.Copy(&UserResp, &user)
	if err != nil {
		return utils.BadRequest(ctx, "internal Server Error", err.Error())
	}
	return utils.Success(ctx, "User retrieved successfully", UserResp)
}
