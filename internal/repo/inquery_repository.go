package repo

import (
	"encoding/json"
	"os"

	"github.com/gin-gonic/gin"
)

type inqueryRepository struct {
}

type InqueryRepository interface {
	GetQuestionnaire(c *gin.Context) ([]map[string]interface{}, error)
}

func NewInqueryRepository() InqueryRepository {
	return &inqueryRepository{}

}

func (r *inqueryRepository) GetQuestionnaire(c *gin.Context) ([]map[string]interface{}, error) {
	// Placeholder for database interaction to retrieve insurance questionnaire data
	questionnaire_path := "../../shared/questionnaires.json"
	file, err := os.ReadFile(questionnaire_path)
	if err != nil {
		return nil, err
	}

	var data []map[string]interface{}
	err = json.Unmarshal(file, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
