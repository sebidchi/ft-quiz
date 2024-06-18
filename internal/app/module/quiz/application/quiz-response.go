package application

import "github.com/sebidchi/ft-quiz/internal/app/module/quiz/domain"

// QuizResponse is DTO that contains the questions of a quiz
// and the answers of the questions
type QuizResponse struct {
	Questions []QuestionResponse
}

type QuestionResponse struct {
	Id           string
	QuestionText string
	Options      []OptionResponse
	Answer       string
}

type OptionResponse struct {
	Id   string
	Text string
}

// NewQuizResponseFromQuiz creates a QuizResponse from a Quiz
func NewQuizResponseFromQuiz(quiz *domain.Quiz) *QuizResponse {
	var questions []QuestionResponse

	for _, question := range quiz.Questions() {
		var options []OptionResponse

		for _, option := range question.Options() {
			options = append(options, OptionResponse{
				Id:   option.Id(),
				Text: option.Text(),
			})
		}

		questions = append(questions, QuestionResponse{
			Id:           question.Id(),
			QuestionText: question.Text(),
			Options:      options,
			Answer:       question.Answer(),
		})
	}

	return &QuizResponse{
		Questions: questions,
	}
}
