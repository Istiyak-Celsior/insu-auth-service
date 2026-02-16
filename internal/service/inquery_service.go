package service

import (
	"InsuranceChatWS/internal/repo"

	"github.com/gin-gonic/gin"
)

type inqueryService struct {
	inqueryRepository repo.InqueryRepository
}

type InqueryService interface {
	GetQuestionnaireService(c *gin.Context) ([]map[string]interface{}, error)
}

func NewInqueryService(inqueryRepository repo.InqueryRepository) InqueryService {
	return &inqueryService{
		inqueryRepository: inqueryRepository,
	}
}

func (s *inqueryService) GetQuestionnaireService(c *gin.Context) ([]map[string]interface{}, error) {
	return s.inqueryRepository.GetQuestionnaire(c)
}
