package users

import (
	users "bctec/internal"
	"github.com/gin-gonic/gin"
	"net/http"
)

type createRequest struct {
	username  string `json:"username" binding:"required"`
	email     string `json:"email" binding:"required"`
	firstName string `json:"firstName" binding:"required"`
	lastName  string `json:"lastName" binding:"required"`
	website   string `json:"website"`
	password  string `json:"password" binding:"required"`
	role      string `json:"role" binding:"required"`
}

func CreateHandler(usersRepository users.UserRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req createRequest

		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		if len(req.website) > 0 {
			user := users.NewUser(req.username, req.email, req.firstName, req.lastName, req.website, req.password, req.role)

			if err := usersRepository.Save(ctx, user); err != nil {
				ctx.JSON(http.StatusInternalServerError, err.Error())
				return
			}

			ctx.Status(http.StatusCreated)
		} else {
			user := users.NewUser(req.username, req.email, req.firstName, req.lastName, "", req.password, req.role)

			if err := usersRepository.Save(ctx, user); err != nil {
				ctx.JSON(http.StatusInternalServerError, err.Error())
				return
			}

			ctx.Status(http.StatusCreated)
		}
	}
}
