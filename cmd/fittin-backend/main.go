package main

import (
	"database/sql"
	"log"

	"github.com/Dmitriy-M1319/fittin-backend/internal/mmpi/api"
	"github.com/Dmitriy-M1319/fittin-backend/internal/mmpi/db"
	"github.com/Dmitriy-M1319/fittin-backend/internal/mmpi/services"
)

func main() {

	database, err := sql.Open("sqlite3", "./test_results.db")
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	repository := db.NewSQLiteTestResultRepository(database)
	resultService := services.NewTestResultService(repository)
	calcService := services.NewTestCalculationService()
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
