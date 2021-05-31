package models

// Error ...
type Error struct {
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
}

//Response ...
type Response struct {
	Code  int         `json:"status"`
	Error *Error      `json:"error"`
	Data  interface{} `json:"data"`
	Count interface{} `json:"count,omitempty"`
}
