// Package constants defines reusable response messages for the API.
package constants

const (
	MsgIncidentCreated = "Incident created successfully."
	MsgInvalidPayload  = "Invalid request payload."
	MsgInternalError   = "Internal server error."
)

type DBConnCtxKey string

const (
	DB_CONN_CTX_KEY DBConnCtxKey = "dbConn"
)
