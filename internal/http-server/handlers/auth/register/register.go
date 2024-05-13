package register

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	ID int64 `json:"user_id"`
}

type RegisterRequest struct {
	Email    string `json:"email"`    // Email of the user to register.
	Password string `json:"password"` // Password of the user to register.
}

type UserRegisterer interface {
	Register(email string, password string) (int64, error)
}

func MakeGetHandlerFunc(log *slog.Logger, registerer UserRegisterer) gin.HandlerFunc {
	const op = "http-server.handlers.auth.register.MakeGetHandlerFunc"

	log = log.With(
		slog.String("op", op),
	)

	return func(c *gin.Context) {

		var req RegisterRequest

		if err := c.BindJSON(&req); err != nil {
			log.Error("err: ", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		id, err := registerer.Register(req.Email, req.Password)
		if err != nil {
			log.Error("error while registration")

			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"user_id": id,
		})
	}
}
