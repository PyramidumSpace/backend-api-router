package get

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pyramidum-space/protos/gen/go/tasks"
)

type Task struct {
	D                []byte               `json:"id"`
	Header           string               `json:"header"`
	Text             string               `json:"text"`
	ExternalImages   []string             `json:"external_images"`
	Deadline         time.Time            `json:"deadline"`
	ProgressStatus   tasks.ProgressStatus `json:"progress_status"`
	IsUrgent         bool                 `json:"is_urgent"`
	IsImportant      bool                 `json:"is_important"`
	OwnerID          int32                `json:"owner_id"`
	ParentID         []byte               `json:"parent_id"`
	PossibleDeadline time.Time            `json:"possible_deadline"`
	Weight           int32                `json:"weight"`
}
type Response struct {
	Tasks []Task `json:"tasks"`
}
type TaskGetter interface {
	Get(taskId int32) ([]*tasks.Task, error)
}

func MakeGetHandlerFunc(log *slog.Logger, getter TaskGetter) gin.HandlerFunc {
	const op = "http-server.handlers.tasks.get.MakeGetHandlerFunc"

	log = log.With(
		slog.String("op", op),
	)

	return func(c *gin.Context) {

		userIdStr := c.Query("user_id")
		if userIdStr == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "The user_id parameter was not passed",
			})
			return
		}

		userIdInt, err := strconv.Atoi(userIdStr)
		if err != nil {
			fmt.Println("Ошибка:", err)
			return
		}

		task_list, err := getter.Get(int32(userIdInt))
		if err != nil {
			log.Error("error while registration")

			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"tasks": task_list,
		})
	}
}
