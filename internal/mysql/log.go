package mysql

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/Team-Missing-Semicolons/cyberhackathon-2024-backend/internal/domain"
)

// getLogsSql is the base query for selecting all logs from the database.
const getLogsSql = `
SELECT id, timestamp, log_type, data FROM logs
`

// insertLog is a SQL statement that inserts a new entry into logs table.
const insertLogSql = `
INSERT INTO logs (timestamp, log_type, data, unit_id) values (?, ?, ?, ?)
`

// LogDataStore represents a log data store for mysql database.
type LogDataStore struct {
	db *sql.DB
}

// NewLogDataStore creates a new instance of a data store.
func NewLogDataStore(db *sql.DB) *LogDataStore {
	return &LogDataStore{
		db: db,
	}
}

// InsertLog inserts a log to the database.
func (ds *LogDataStore) InsertLog(ctx context.Context, logInsertPayload *domain.LogInsertPayload) error {
	rawData, err := json.Marshal(logInsertPayload.Data)
	if err != nil {
		return fmt.Errorf("failed to marshal log insert payload: %w", err)
	}

	_, err = ds.db.ExecContext(
		ctx,
		insertLogSql,
		logInsertPayload.Timestamp,
		logInsertPayload.Type,
		rawData,
		logInsertPayload.UnitID,
	)
	if err != nil {
		return fmt.Errorf("failed to exec: %w", err)
	}

	return nil
}

// GetLog selects all logs from the database.
func (ds *LogDataStore) GetLog(ctx context.Context) ([]*domain.Log, error) {
	logs := make([]*domain.Log, 0)

	rows, err := ds.db.QueryContext(ctx, getLogsSql)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var logEntry domain.Log
		var rawData []byte

		if err := rows.Scan(&logEntry.Id, &logEntry.Timestamp, &logEntry.Type, &rawData); err != nil {
			fmt.Println("ERROR: failed to scan log row")
			continue
		}

		if err := json.Unmarshal(rawData, &logEntry.Data); err != nil {
			fmt.Println("failed to unmarshal log data")
			continue
		}

		logs = append(logs, &logEntry)
	}

	return logs, nil
}
