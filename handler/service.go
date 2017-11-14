package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

// StringService provides operations on strings.
type Service interface {
	Count(c *gin.Context)
	Ping(c *gin.Context)
}
type Str struct {
	S string `json:"s"`
}
type stringService struct{}

func NewService() Service {
	return &stringService{}
}
func (s *stringService) Ping(c *gin.Context) {
	glog.Info("get ping")
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
func (s *stringService) Count(c *gin.Context) {
	var str Str
	if c.Bind(&str) == nil {
		glog.Info("get string:", str.S)
		c.JSON(200, gin.H{"count": len(str.S)})
	}
}
