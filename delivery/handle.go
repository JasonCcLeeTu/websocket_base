package handle

import "github.com/gin-gonic/gin"

func Router(engine *gin.Engine) {
	rounter := engine.Group("/ping")
	rounter.GET("", WsPing)
}
