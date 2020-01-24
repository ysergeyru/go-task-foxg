package server

// ok is a standard ACK response when requesting [API]/ok with a GET verb
var ok = map[string]interface{}{"ok": true}

// pong is a standard ACK response when requesting [API]/ping with a GET verb
var pong = map[string]interface{}{"pong": true}

const (
	// Supported HTTP Methods for the server
	POST    = "POST"
	GET     = "GET"
	PUT     = "PUT"
	OPTIONS = "OPTIONS"
	DELETE  = "DELETE"
)

// Headers
const (
	XUserID = "X-User"
)
