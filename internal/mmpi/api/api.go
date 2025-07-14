package api

import (
	"context"

	_ "github.com/Dmitriy-M1319/fittin-backend/docs" // важно добавить этот импорт
	"github.com/Dmitriy-M1319/fittin-backend/internal/mmpi/models"
	"github.com/Dmitriy-M1319/fittin-backend/internal/mmpi/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

// @title MMPI Test API
// @version 1.0
// @description API for MMPI psychological test administration
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email support@fittin.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /

// AttemptRequest представляет запрос с UUID попытки
type AttemptRequest struct {
	Uuid string `json:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

// AddNewAnswerRequest представляет запрос на добавление ответа
type AddNewAnswerRequest struct {
	Uuid   string        `json:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	Answer models.Answer `json:"answer"`
}

type MMPITestApi struct {
	app                *fiber.App
	attemptService     *services.TestAttemptService
	resultService      *services.TestResultService
	calculationService *services.TestCalculationService
	questionService    *services.QuestionService
}

func NewMMPITestApi(attSrv *services.TestAttemptService,
	resSrv *services.TestResultService,
	calcSrv *services.TestCalculationService,
	quesSrv *services.QuestionService) *MMPITestApi {
	return &MMPITestApi{
		app:                fiber.New(),
		attemptService:     attSrv,
		resultService:      resSrv,
		calculationService: calcSrv,
		questionService:    quesSrv,
	}
}

func (a *MMPITestApi) RegisterServices(addr string) {
	// Добавляем маршрут для Swagger UI
	a.app.Get("/docs/*", swagger.HandlerDefault)

	a.app.Get("/questions", a.HandleGetQuestions)
	a.app.Post("/attempt", a.HandleCreateNewAttempt)
	a.app.Post("/answer", a.HandleAddNewAnswer)
	a.app.Post("/calculate", a.HandleCalculateResult)
	a.app.Listen(addr)
}

// HandleGetQuestions godoc
// @Summary Получить все вопросы теста
// @Description Возвращает список всех вопросов MMPI теста
// @Tags questions
// @Accept json
// @Produce json
// @Success 200 {array} models.Question
// @Failure 404 {object} map[string]string
// @Router /questions [get]
func (a *MMPITestApi) HandleGetQuestions(c *fiber.Ctx) error {
	questions := a.questionService.GetQuestions()
	if questions == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Can not find questions",
		})
	}
	return c.JSON(questions)
}

// HandleCreateNewAttempt godoc
// @Summary Создать новую попытку тестирования
// @Description Создает новую запись для хранения ответов на тест
// @Tags attempt
// @Accept json
// @Produce json
// @Param request body AttemptRequest true "Запрос на создание попытки"
// @Success 201
// @Failure 400 {object} map[string]string
// @Router /attempt [post]
func (a *MMPITestApi) HandleCreateNewAttempt(c *fiber.Ctx) error {
	var req AttemptRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	a.attemptService.CreateNewAttempt(req.Uuid)
	return c.SendStatus(fiber.StatusCreated)
}

// HandleAddNewAnswer godoc
// @Summary Добавить ответ на вопрос
// @Description Сохраняет ответ на конкретный вопрос теста
// @Tags attempt
// @Accept json
// @Produce json
// @Param request body AddNewAnswerRequest true "Запрос с ответом"
// @Success 200
// @Failure 400 {object} map[string]string
// @Router /answer [post]
func (a *MMPITestApi) HandleAddNewAnswer(c *fiber.Ctx) error {
	var req AddNewAnswerRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	err := a.attemptService.AddAnswer(req.Uuid, req.Answer)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.SendStatus(fiber.StatusOK)
}

// HandleCalculateResult godoc
// @Summary Рассчитать результаты теста
// @Description Вычисляет результаты теста на основе предоставленных ответов
// @Tags results
// @Accept json
// @Produce json
// @Param request body AttemptRequest true "Запрос с UUID попытки"
// @Success 200 {object} models.TestResult
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /calculate [post]
func (a *MMPITestApi) HandleCalculateResult(c *fiber.Ctx) error {
	var req AttemptRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	attempt := a.attemptService.GetAttemptByUUID(req.Uuid)
	if attempt == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Attempt not found",
		})
	}

	result, err := a.calculationService.Calculate(attempt)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Can not calculate test result",
		})
	}

	err = a.resultService.AddNewResult(context.Background(), result)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Can not save test result",
		})
	}

	return c.JSON(result)
}
