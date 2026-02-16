package handler

import (
	"InsuranceChatWS/internal/service"

	"github.com/gin-gonic/gin"
)

type inqueryHandler struct {
	inqueryService service.InqueryService
}

type InqueryHandler interface {
	GetQuestionnaire(c *gin.Context)
}

func NewInqueryHandler(inqueryService service.InqueryService) InqueryHandler {
	return &inqueryHandler{
		inqueryService: inqueryService,
	}
}

func (h *inqueryHandler) GetQuestionnaire(c *gin.Context) {
	data, error := h.inqueryService.GetQuestionnaireService(c)
	if error != nil {
		c.JSON(500, gin.H{
			"HttpStatusCode": 500,
			"ResponseBody":   nil,
			"IsSuccess":      false,
			"Message":        "Failed to retrieve questionnaire data",
		})
		return
	}

	c.JSON(200, gin.H{
		"HttpStatusCode": 200,
		"ResponseBody":   data,
		"IsSuccess":      true,
		"Message":        "Questionnaire data retrieved successfully",
	})
}
