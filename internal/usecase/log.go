package usecase

import (
	"context"
	"fmt"

	"github.com/Team-Missing-Semicolons/cyberhackathon-2024-backend/internal/domain"
	"github.com/Team-Missing-Semicolons/cyberhackathon-2024-backend/internal/mysql"
)

// LogUseCase represents a structure storing functionality related to log aggregation.
type LogUseCase struct {
	logDataStore *mysql.LogDataStore
}

// NewLogUseCase creates a new instance of LogUseCase.
func NewLogUseCase(logDataStore *mysql.LogDataStore) *LogUseCase {
	return &LogUseCase{
		logDataStore: logDataStore,
	}
}

// InsertLog processes the log and stores it in a database for further processing.
func (uc *LogUseCase) InsertLog(ctx context.Context, logInsertPayload *domain.LogInsertPayload) error {
	if err := uc.logDataStore.InsertLog(ctx, logInsertPayload); err != nil {
		return fmt.Errorf("failed to store the log: %w: %w", ErrDataStoreFailed, err)
	}
	return nil
}

// GetLogs is a method for getting logs.
func (uc *LogUseCase) GetLogs(ctx context.Context) ([]*domain.Log, error) {
	logs, err := uc.logDataStore.GetLog(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch logs: %w", err)
	}
	return logs, nil
}
