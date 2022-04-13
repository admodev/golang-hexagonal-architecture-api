package users

import (
	"bctec/cmd/api/environment"
	users "bctec/internal"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
)

type JWTClaims struct {
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

type createRequest struct {
	Token     string `json:"token"`
	Username  string `json:"username" binding:"required"`
	Email     string `json:"email" binding:"required"`
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
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
	secret := environment.SECRET
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

	fmt.Println("Signed token: " + signedJwtoken)

	return signedJwtoken
}

func CreateHandler(usersRepository users.UserRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req createRequest

		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		jwtUserToken := CreateJWT(req.Email, req.Role)

		if len(req.Website) > 0 {
			user := users.NewUser(jwtUserToken, req.Username, req.Email, req.FirstName, req.LastName, req.Website, req.Password, req.Role)

			if err := usersRepository.Save(ctx, user); err != nil {
				ctx.JSON(http.StatusInternalServerError, err.Error())
				return
			}

			ctx.Status(http.StatusCreated)
		} else {
			user := users.NewUser(jwtUserToken, req.Username, req.Email, req.FirstName, req.LastName, "none", req.Password, req.Role)

			if err := usersRepository.Save(ctx, user); err != nil {
				ctx.JSON(http.StatusInternalServerError, err.Error())
				return
			}

			ctx.Status(http.StatusCreated)
		}
	}
}
