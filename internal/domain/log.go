package domain

import "context"

// LogInsertPayload represents the body payload processed by the server.
type LogInsertPayload struct {
	Type      string      `json:"type"`      // Stores the log Type
	Timestamp int64       `json:"timestamp"` // Timestamp of when the log was produced.
	Data      interface{} `json:"data"`      // Data stored in the log.
}

// Log represents a log struct.
type Log struct {
	Id        string      `json:"id"`        // Log id.
	Type      string      `json:"type"`      // Stores the log Type
	Timestamp int64       `json:"timestamp"` // Timestamp of when the log was produced.
	Data      interface{} `json:"data"`      // Data stored in the log.
}

// LogDataSource defines methods that must be implemented by LogDataSource.
type LogDataSource interface {
	// InsertLog inserts a log to the a database.
	InsertLog(ctx context.Context, logType string, data []byte) error
}
