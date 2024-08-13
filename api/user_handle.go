package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/unsuman/hotel-reservation.git/db"
	"github.com/unsuman/hotel-reservation.git/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserHandler struct {
	store db.Store
}

func NewUserHandler(store *db.Store) *UserHandler {
	return &UserHandler{
		store: *store,
	}
}

func (h *UserHandler) HandleUpdateUser(c *fiber.Ctx) error {
	var (
		update types.UpdateUserParams
	)
	id := c.Params("id")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": oid}

	if err := c.BodyParser(&update); err != nil {
		return err
	}

	if err := h.store.UserStore.UpdateUser(c.Context(), filter, update); err != nil {
		return err
	}
	return c.JSON(map[string]string{"updated": id})
}

func (h *UserHandler) HandleDeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := h.store.UserStore.DeleteUser(c.Context(), id); err != nil {
		return err
	}

	return c.JSON(map[string]string{"Status": "Deleted " + id})
}

func (h *UserHandler) HandlePostUser(c *fiber.Ctx) error {
	var params types.CreateUsersParams

	if err := c.BodyParser(&params); err != nil {
		fmt.Println(err)
		return err
	}

	if err := params.Validate(); len(err) > 0 {
		return c.JSON(err)
	}

	user, err := types.NewUserFromParams(params)
	if err != nil {
		return err
	}

	insertedUser, err := h.store.UserStore.InsertUser(c.Context(), user)
	if err != nil {
		return err
	}

	return c.JSON(insertedUser)
}

func (h *UserHandler) HandleGetUsers(c *fiber.Ctx) error {
	users, err := h.store.UserStore.GetUsers(c.Context())
	if err != nil {
		return err
	}
	return c.JSON(users)
}

func (h *UserHandler) HandleGetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := h.store.UserStore.GetUserByID(c.Context(), id)
	if err != nil {
		return err
	}

	return c.JSON(user)
}
