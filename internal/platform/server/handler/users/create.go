package users

import (
	"bctec/cmd/api/environment"
	users "bctec/internal"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
	"net/http"
)

var UserJwt string

type JWTClaims struct {
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

type createRequest struct {
	gorm.Model
	Id        int64  `gorm:"primary_key;auto_increment;not_null"`
	Token     string `json:"token"`
	Username  string `json:"username" binding:"required"`
	Email     string `json:"email" binding:"required"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Website   string `json:"website"`
	Password  string `json:"password" binding:"required"`
	Role      string `json:"role" binding:"required"`
}

func CreateJWT(email, role string) string {
	claims := JWTClaims{
		email,
		role,
		jwt.StandardClaims{
			ExpiresAt: 15000,
			Issuer:    "BCTEC",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	unparsedSecret := environment.SECRET
	secret := []byte(unparsedSecret)
	signedJwtoken, err := token.SignedString(secret)

	if err != nil {
		jwtError := fmt.Sprintf("Error signing JSONWebToken! %v", err.Error())

		return jwtError
	}

	token, err = jwt.Parse(signedJwtoken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signin method: %v", token.Header["alg"])
		}

		return secret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["email"], claims["role"])
	} else {
		fmt.Sprintf(err.Error())
	}

	UserJwt = signedJwtoken

	return signedJwtoken
}

func CreateHandler(usersRepository users.UserRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req createRequest

		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		CreateJWT(req.Email, req.Role)

		if len(req.Website) > 0 {
			user := users.NewUser(UserJwt, req.Username, req.Email, req.FirstName, req.LastName, req.Website, req.Password, req.Role)

			if err := usersRepository.Save(ctx, user); err != nil {
				ctx.JSON(http.StatusInternalServerError, err.Error())
				return
			}

			ctx.Status(http.StatusCreated)
		} else {
			user := users.NewUser(UserJwt, req.Username, req.Email, req.FirstName, req.LastName, "none", req.Password, req.Role)

			if err := usersRepository.Save(ctx, user); err != nil {
				ctx.JSON(http.StatusInternalServerError, err.Error())
				return
			}

			ctx.Status(http.StatusCreated)
		}
	}
}
