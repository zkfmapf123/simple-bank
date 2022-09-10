package helpers

import "github.com/gin-gonic/gin"

const (
	ALREADY_EXISTS_OWNER = "already-exists-user"
	NOT_EXISTS_OWNER = "not-exists-user"
)

func ErrorReason(err error, msg string) gin.H {

	if err == nil {
		return gin.H{"status" : "error", "message" : msg}
	}

	return gin.H{"error": err.Error(), "message" : msg}
}