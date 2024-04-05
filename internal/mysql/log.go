package mysql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Team-Missing-Semicolons/cyberhackathon-2024-backend/internal/domain"
)

// getLogsSql is the base query for selecting all logs from the database.
const getLogsSql = `
SELECT id, timestamp, log_type, data FROM logs
`

// insertLog is a SQL statement that inserts a new entry into logs table.
const insertLogSql = `
INSERT INTO logs (timestamp, log_type, data) values (?, ?, ?)
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
func (ds *LogDataStore) InsertLog(ctx context.Context, timestamp int64, logType string, data []byte) error {
	if _, err := ds.db.ExecContext(ctx, insertLogSql, timestamp, logType, data); err != nil {
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
		if err := rows.Scan(&logEntry.Id, &logEntry.Timestamp, &logEntry.Type, &logEntry.Data); err != nil {
			fmt.Println("ERROR: failed to scan log row")
			continue
		}
		logs = append(logs, &logEntry)
	}

	return logs, nil
}
