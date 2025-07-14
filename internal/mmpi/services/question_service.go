package services

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/Dmitriy-M1319/fittin-backend/internal/mmpi/models"
)

type QuestionService struct {
	questions []*models.Question
}

func NewQuestionService() *QuestionService {
	return &QuestionService{questions: nil}
}

func (s *QuestionService) LoadQuestions(sheetID, sheetName string) error {
	url := fmt.Sprintf("https://docs.google.com/spreadsheets/d/%s/gviz/tq?tqx=out:csv&sheet=\"%s\"", sheetID, url.QueryEscape(sheetName))

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("HTTP request failed: %v", err)
	}
	defer resp.Body.Close()

	log.Print(resp)

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	reader := csv.NewReader(resp.Body)
	lineNumber := 0

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("CSV read error: %v", err)
		}

		if len(record) == 0 {
			continue
		}

		s.questions = append(s.questions, &models.Question{
			Number: lineNumber,
			Text:   record[0],
		})
		lineNumber++
	}

	return nil
}

func (s *QuestionService) GetQuestions() []*models.Question {
	return s.questions
}
