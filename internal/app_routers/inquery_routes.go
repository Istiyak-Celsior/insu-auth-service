package approuters

import (
	"InsuranceChatWS/internal/configuration"

	"github.com/gin-gonic/gin"
)

func InqueryRouters(router *gin.Engine, container *configuration.Container) {
	inqueryRouter := router.Group("/inc/api")
	{
		inqueryRouter.GET("/questionaries/start", container.InqueryHandler.GetQuestionnaire)
	}
}
