package handler_test

import (
	"net/http"
	"goexample/handler"
	"testing"

	"github.com/appleboy/gofight"
	"github.com/buger/jsonparser"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var s = handler.NewService()

func GinEngine() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/ping", s.Ping)
	r.POST("/count", s.Count)
	return r
}
func TestStringService_Ping(t *testing.T) {
	r := gofight.New()
	r.GET("/ping").
		SetDebug(true).
		Run(GinEngine(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusOK, r.Code)
		})
}

func TestStringService_Count(t *testing.T) {
	r := gofight.New()
	r.POST("/count").
		SetJSON(gofight.D{
			"S": "dasffff",
		}).
		SetDebug(true).
		Run(GinEngine(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusOK, r.Code)
			t.Log(r.Body.String())
			data := r.Body.Bytes()
			l, _ := jsonparser.GetInt(data, "count")
			assert.Equal(t, len("dsfasfd"), int(l))
		})
}
