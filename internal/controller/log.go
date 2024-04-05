package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Team-Missing-Semicolons/cyberhackathon-2024-backend/internal/domain"
	"github.com/Team-Missing-Semicolons/cyberhackathon-2024-backend/internal/usecase"
)

type LogController struct {
	logUseCase *usecase.LogUseCase
}

func NewLogController(logUseCase *usecase.LogUseCase) *LogController {
	return &LogController{
		logUseCase: logUseCase,
	}
}

// HandleInsertLog is an API handler for inserting a new log into the database.
func (c *LogController) HandleInsertLog(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var logInsertPayload domain.LogInsertPayload
	if err := json.NewDecoder(r.Body).Decode(&logInsertPayload); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := c.logUseCase.InsertLog(ctx, &logInsertPayload); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	return
}

// HandleGetLogs responds with a list of logs.
func (c *LogController) HandleGetLogs(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	logs, err := c.logUseCase.GetLogs(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	responseBody, err := json.Marshal(logs)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseBody)
}
