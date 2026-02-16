package approuters

import (
	"InsuranceChatWS/internal/configuration"
	"InsuranceChatWS/internal/handler"
	"InsuranceChatWS/internal/hub"

	"github.com/gin-gonic/gin"
)

// MonitorRouters sets up monitoring API routes
func MonitorRouters(router *gin.Engine, container *configuration.Container) {
	// Create monitor service with hub reference
	monitorService := hub.NewMonitorService(container.Hub)

	// Create monitor handler
	monitorHandler := handler.NewMonitorHandler(monitorService)

	// Monitor API group
	monitorGroup := router.Group("/chat/api/monitor")
	{
		// GET /api/monitor/stats - Get hub statistics
		monitorGroup.GET("/stats", monitorHandler.GetHubStats)
	}
}
