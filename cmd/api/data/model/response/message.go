package response

// MessageResponse used to respond to the client with a simple JSON message
type MessageResponse struct {
	Message string `json:"message"`
}
