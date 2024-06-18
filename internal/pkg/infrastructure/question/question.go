package question

type AnswersPayload struct {
	Username string       `json:"username"`
	Answers  []UserAnswer `json:"answers"`
}
type Quiz struct {
	ID        string     `json:"id"`
	Title     string     `json:"title"`
	Questions []Question `json:"questions"`
}

// Question represents a quiz question
type Question struct {
	ID             string    `json:"question"`
	Question       string    `json:"questions"`
	Options        []*Option `json:"options"`
	AnswerOptionId string    `json:"answer"`
}

type Option struct {
	OptionId string `json:"id"`
	Text     string `json:"text"`
}

// UserAnswer represents an answer submitted by a user
type UserAnswer struct {
	QuestionID string `json:"question_id"`
	Answer     string `json:"answer"`
}

// Score represents a user's score
type Score struct {
	Correct int `json:"correct"`
	Total   int `json:"total"`
}

type UserResults struct {
	UserId     string  `json:"user_id"`
	Total      int     `json:"total"`
	Percentage float64 `json:"percentage"`
	BetterThan float64 `json:"better_than"`
	TotalUsers int     `json:"total_users"`
}
