package login

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	ID int64 `json:"user_id"`
}

type LoginRequest struct {
	Email    string `json:"email"`    // Email of the user to login.
	Password string `json:"password"` // Password of the user to login.
}

type UserLoginer interface {
	Login(email string, password string) (int64, error)
}

func MakeGetHandlerFunc(log *slog.Logger, loginer UserLoginer) gin.HandlerFunc {
	const op = "http-server.handlers.auth.login.MakeGetHandlerFunc"

	log = log.With(
		slog.String("op", op),
	)

	return func(c *gin.Context) {

		var req LoginRequest

		if err := c.BindJSON(&req); err != nil {
			log.Error("err: ", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		id, err := loginer.Login(req.Email, req.Password)
		if err != nil {
			log.Error("error while login")

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
