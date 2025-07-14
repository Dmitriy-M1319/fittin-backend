package models

type Question struct {
	Number int    `json:"number"`
	Text   string `json:"text"`
}

// Тип ответа
const (
	RIGHT = iota
	WRONG
	NO_ANSWER
)

type Answer struct {
	QuestionNumber int `json:"question_number"`
	AnswerVariant  int `json:"answer_variant"`
}

type TestAttempt struct {
	Uuid    string
	Answers []Answer
}

type ScalingResult struct {
	Scale string `json:"scale"`
	Value int    `json:"value"`
}

type TestResult struct {
	Uuid   string          `json:"uuid"`
	Scales []ScalingResult `json:"scales"`
	Info   string          `json:"information"`
}
