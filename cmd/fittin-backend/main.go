package main

import (
	"database/sql"
	"errors"
	"log"
	"os"

	"github.com/Dmitriy-M1319/fittin-backend/internal/mmpi/api"
	"github.com/Dmitriy-M1319/fittin-backend/internal/mmpi/config"
	"github.com/Dmitriy-M1319/fittin-backend/internal/mmpi/db"
	"github.com/Dmitriy-M1319/fittin-backend/internal/mmpi/llm"
	"github.com/Dmitriy-M1319/fittin-backend/internal/mmpi/services"
)

func main() {

	if err := config.ReadConfigYML("config.yml"); err != nil {
		log.Fatalf("Failed init configuration")
	}

	cfg := config.GetConfigInstance()

	if _, err := os.Stat("./test_results.db"); errors.Is(err, os.ErrNotExist) {
		os.Create("./test_results.db")
	}

	database, err := sql.Open("sqlite3", "./test_results.db")
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	llmService := llm.NewLLMClient(cfg)
	repository := db.NewSQLiteTestResultRepository(database)
	resultService := services.NewTestResultService(repository)
	calcService := services.NewTestCalculationService(llmService)
	questionService := services.NewQuestionService()
	attService := services.NewTestAttemptService()

	sheetID := "19Qr5HKKAsRFbOonQqoV3jOLcUwfHKcnkaVzhI-s_OPw"
	sheetName := "Мужские вопросы"
	err = questionService.LoadQuestions(sheetID, sheetName)
	if err != nil {
		log.Fatal(err.Error())
	}

	restApp := api.NewMMPITestApi(attService, resultService, calcService, questionService)
	restApp.RegisterServices(":8050")
}
