package service

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine) {

	r.GET("/miners", GetCacheData)
	r.POST("/push", ReceivedCacheData)
}
