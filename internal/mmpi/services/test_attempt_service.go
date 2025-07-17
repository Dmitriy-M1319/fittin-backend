package services

import (
	"fmt"
	"log"

	"github.com/Dmitriy-M1319/fittin-backend/internal/mmpi/models"
)

type TestAttemptService struct {
	attempts map[string]*models.TestAttempt
}

func NewTestAttemptService() *TestAttemptService {
	return &TestAttemptService{attempts: make(map[string]*models.TestAttempt, 0)}
}

func (s *TestAttemptService) AddAnswer(uuid string, answer models.Answer) error {
	if s.attempts[uuid] != nil {
		att := s.attempts[uuid]
		att.Answers = append(att.Answers, answer)
		return nil
	} else {
		return fmt.Errorf("invalid uuid")
	}
}

func (s *TestAttemptService) SetAllAnswers(uuid string, variant int) {
	for i := 1; i <= 566; i += 1 {
		s.AddAnswer(uuid, models.Answer{QuestionNumber: i, AnswerVariant: variant})
	}
	log.Println(len(s.attempts[uuid].Answers))
}

func (s *TestAttemptService) CreateNewAttempt(uuid string) {
	s.attempts[uuid] = &models.TestAttempt{Uuid: uuid}
}

func (s *TestAttemptService) GetAttemptByUUID(uuid string) *models.TestAttempt {
	return s.attempts[uuid]
}
