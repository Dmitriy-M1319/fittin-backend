package services

import (
	"errors"
	"fmt"

	"github.com/Dmitriy-M1319/fittin-backend/internal/mmpi/models"
)

type scale struct {
	True  []int `json:"true"`
	False []int `json:"false"`
}

var clinicalScales = map[string]scale{
	"(Hs)": { // Ипохондрия
		True:  []int{23, 29, 43, 62, 72, 108, 114, 125, 161, 189, 273},
		False: []int{2, 3, 7, 9, 18, 51, 55, 63, 68, 103, 140, 153, 155, 163, 175, 188, 190, 192, 230, 243, 274, 281},
	},
	"(D)": { // Депрессия
		True:  []int{5, 13, 23, 32, 41, 43, 52, 67, 86, 104, 130, 138, 142, 158, 159, 182, 189, 193, 236, 259},
		False: []int{2, 8, 9, 18, 30, 36, 39, 46, 51, 57, 58, 64, 80, 88, 89, 95, 98, 107, 122, 131, 145, 152, 153, 154, 155, 160, 178, 191, 207, 208, 238, 241, 242, 248, 263, 270, 271, 272, 285, 296},
	},
	"(Hy)": { // Истерия
		True:  []int{10, 23, 32, 43, 44, 47, 76, 114, 179, 186, 189, 238},
		False: []int{2, 3, 6, 7, 8, 9, 12, 26, 30, 51, 55, 71, 89, 93, 103, 107, 109, 124, 128, 129, 136, 137, 141, 147, 153, 160, 162, 163, 170, 172, 174, 175, 180, 188, 190, 192, 201, 213, 230, 234, 243, 265, 267, 274, 279, 289, 292},
	},
	"(Pd)": { // Психопатия
		True:  []int{16, 21, 24, 32, 33, 35, 36, 42, 61, 67, 84, 94, 102, 106, 110, 118, 127, 215, 216, 224, 239, 244, 245, 284},
		False: []int{8, 20, 37, 82, 91, 96, 107, 134, 137, 141, 155, 170, 171, 173, 180, 183, 201, 231, 235, 237, 248, 267, 287, 289, 294, 296},
	},
	"(Mf)": { // Мужественность/Женственность (разные для мужчин и женщин)
		True:  []int{4, 25, 70, 74, 77, 78, 87, 92, 126, 132, 133, 134, 140, 149, 187, 203, 204, 217, 226, 239, 261, 278, 282, 295, 299},
		False: []int{1, 19, 26, 69, 79, 80, 81, 89, 99, 112, 115, 116, 117, 120, 144, 176, 179, 198, 213, 219, 221, 223, 229, 231, 249, 254, 260, 262, 264, 280, 283, 297, 300},
	},
	"(Pa)": { // Паранойя
		True:  []int{15, 16, 22, 24, 27, 35, 110, 121, 123, 127, 151, 157, 158, 202, 275, 284, 291, 293, 299, 305, 317, 338, 341, 364, 365},
		False: []int{93, 107, 109, 111, 117, 124, 268, 281, 294, 313, 316, 319, 327, 347, 348},
	},
	"(Pt)": { // Психастения
		True:  []int{10, 15, 22, 32, 41, 67, 76, 86, 94, 102, 106, 142, 159, 182, 189, 217, 238, 266, 301, 304, 305, 317, 321, 336, 337, 340, 342, 343, 344, 346, 349, 351, 352, 356, 357, 359, 360, 361},
		False: []int{3, 8, 36, 122, 152, 164, 178, 329, 353},
	},
	"(Sc)": { // Шизофрения
		True:  []int{15, 16, 21, 22, 24, 32, 33, 35, 38, 40, 41, 47, 52, 76, 97, 104, 121, 156, 157, 159, 168, 179, 182, 194, 202, 210, 212, 238, 241, 251, 259, 266, 273, 282, 291, 297, 301, 303, 305, 307, 312, 320, 324, 325, 332, 334, 335, 339, 341, 345, 349, 350, 352, 354, 355, 356, 360, 363, 364},
		False: []int{8, 17, 20, 37, 65, 103, 119, 177, 178, 187, 192, 196, 220, 276, 281, 306, 309, 322, 330},
	},
	"(Ma)": { // Гипомания
		True:  []int{11, 13, 21, 22, 59, 64, 73, 97, 100, 109, 127, 134, 145, 156, 157, 167, 181, 194, 212, 222, 226, 228, 232, 233, 238, 240, 250, 251, 263, 266, 268, 271, 277, 279, 298},
		False: []int{101, 105, 111, 119, 130, 148, 166, 171, 180, 267, 289},
	},
	"(Si)": { // Социальная интроверсия
		True:  []int{32, 67, 82, 111, 117, 124, 138, 147, 171, 172, 180, 201, 236, 267, 278, 292, 304, 316, 321, 332, 336, 342, 357, 377, 383, 398, 401, 427, 436, 455, 473, 467, 549, 564},
		False: []int{25, 33, 57, 91, 99, 119, 126, 143, 193, 208, 229, 231, 254, 262, 281, 296, 309, 353, 359, 371, 391, 400, 415, 440, 446, 449, 450, 451, 462, 469, 479, 481, 482, 501, 521, 547},
	},
}

type TestCalculationService struct {
}

func NewTestCalculationService() *TestCalculationService {
	return &TestCalculationService{}
}

func (s *TestCalculationService) Calculate(attempt *models.TestAttempt) (*models.TestResult, error) {
	if attempt == nil {
		return nil, errors.New("test attempt cannot be nil")
	}

	if len(attempt.Answers) == 0 {
		return nil, errors.New("no answers provided")
	}

	result := &models.TestResult{
		Uuid: attempt.Uuid,
	}

	// Создаем быстрый доступ к ответам по номеру вопроса
	answerMap := make(map[int]int)
	for _, answer := range attempt.Answers {
		answerMap[answer.QuestionNumber] = answer.AnswerVariant
	}

	// Обрабатываем только клинические шкалы
	for scaleName, scale := range clinicalScales {
		rawScore := 0

		// Подсчет баллов за верные ответы
		for _, qNum := range scale.True {
			if ans, exists := answerMap[qNum]; exists && ans == 0 {
				rawScore++
			}
		}

		// Подсчет баллов за неверные ответы
		for _, qNum := range scale.False {
			if ans, exists := answerMap[qNum]; exists && ans == 1 {
				rawScore++
			}
		}

		// Конвертируем сырые баллы в T-баллы
		tScore := convertToTScore(scaleName, rawScore)

		// Добавляем результат по шкале
		result.Scales = append(result.Scales, models.ScalingResult{
			Scale: scaleName,
			Value: tScore,
		})
	}

	// Анализ общего профиля
	result.Info = analyzeOverallProfile(result.Scales)

	return result, nil
}

// convertToTScore преобразует сырые баллы в T-баллы
func convertToTScore(scale string, rawScore int) int {
	if rawScore < 20 {
		return 30 + rawScore
	}
	return 50 + (rawScore - 20)
}

// analyzeOverallProfile анализирует общий профиль личности
func analyzeOverallProfile(scales []models.ScalingResult) string {
	highScales := make([]string, 0)
	for _, scale := range scales {
		if scale.Value > 70 {
			highScales = append(highScales, scale.Scale)
		}
	}

	if len(highScales) > 3 {
		return "Сложный профиль с множеством повышенных показателей. Требуется дополнительный анализ."
	}

	if len(highScales) == 0 {
		return "Гармоничный профиль без выраженных пиков"
	}

	return fmt.Sprintf("Профиль с повышенными показателями по шкалам: %v", highScales)
}
