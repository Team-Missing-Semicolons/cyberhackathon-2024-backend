package domain

import "context"

// LogInsertPayload represents the body payload processed by the server.
type LogInsertPayload struct {
	Type      string      `json:"type"`
	Timestamp int64       `json:"timestamp"`
	Data      interface{} `json:"data"`
	UnitID    int         `json:"unitId"`
}

// Log represents a log struct.
type Log struct {
	Id        string      `json:"id"`        // Log id.
	Type      string      `json:"type"`      // Stores the log Type
	Timestamp int64       `json:"timestamp"` // Timestamp of when the log was produced.
	Data      interface{} `json:"data"`      // Data stored in the log.
	UnitID    int         `json:"unitId"`    // UnitID represents a unit that reported the log (for example main hall in case of a factory)
}

// LogDataSource defines methods that must be implemented by LogDataSource.
type LogDataSource interface {
	// InsertLog inserts a log to the a database.
	InsertLog(ctx context.Context, logType string, data []byte) error
}
