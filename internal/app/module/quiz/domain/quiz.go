package domain

type Quiz struct {
	id          string
	title       string
	description string
	questions   []Question
}

func NewQuiz(id string, title string, description string, questions []Question) *Quiz {
	return &Quiz{id: id, title: title, description: description, questions: questions}
}

func (q Quiz) Title() string {
	return q.title
}

func (q Quiz) Description() string {
	return q.description
}

func (q Quiz) Questions() []Question {
	return q.questions
}
