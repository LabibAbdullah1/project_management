package controllers

import (
	"ProjectManagement/models"
	"ProjectManagement/services"
	"ProjectManagement/utils"
	"math"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
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
	err = copier.Copy(&UserResp, &user)
	if err != nil {
		return utils.BadRequest(ctx, "internal Server Error", err.Error())
	}
	return utils.Success(ctx, "User retrieved successfully", UserResp)
}

func (c *UserController) GetUsersPageination(ctx *fiber.Ctx) error {
	page, _ := strconv.Atoi(ctx.Query("page", "1"))
	limit, _ := strconv.Atoi(ctx.Query("limit", "10"))
	ofset := (page - 1) * limit

	filter := ctx.Query("filter", "")
	sort := ctx.Query("sort", "created_at desc")

	users, total, err := c.service.GetAllPagination(filter, sort, limit, ofset)
	if err != nil {
		return utils.BadRequest(ctx, "internal Server Error", err.Error())
	}
	var userResp []models.UserResponse
	_ = copier.Copy(&userResp, &users)

	meta := utils.PageinationMeta{
		Page:      page,
		Limit:     limit,
		Total:     int(total),
		TotalPage: int(math.Ceil(float64(total) / float64(limit))),
		Filter:    filter,
		Sort:      sort,
	}

	if total == 0 {
		return utils.NotFoundPageination(ctx, "Data not Found", userResp, meta)
	}
	return utils.SuccessPageination(ctx, "Users retrieved successfully", userResp, meta)
}

func (c *UserController) UpdateUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	publicID, err := uuid.Parse(id)
	if err != nil {
		return utils.BadRequest(ctx, "Invalid User ID Format", err.Error())
	}

	var user models.User
	if err := ctx.BodyParser(&user); err != nil {
		return utils.BadRequest(ctx, "Failed to parse request body", err.Error())
	}
	user.PublicID = publicID

	if err := c.service.Update(&user); err != nil {
		return utils.BadRequest(ctx, "Failed to update user", err.Error())
	}

	userUpdated, err := c.service.GetByPublicID(id)
	if err != nil {
		return utils.InternalServerError(ctx, "Failed to retrieve updated user", err.Error())
	}

	var UserResp models.UserResponse
	err = copier.Copy(&UserResp, &userUpdated)
	if err != nil {
		return utils.InternalServerError(ctx, "Error Parsing Data", err.Error())
	}
	return utils.Success(ctx, "User updated successfully", UserResp)
}

func (c *UserController) DeleteUser(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))
	if err := c.service.Delete(uint(id)); err != nil {
		return utils.InternalServerError(ctx, "Failed to delete user", err.Error())
	}
	return utils.Success(ctx, "User deleted successfully", nil)
}