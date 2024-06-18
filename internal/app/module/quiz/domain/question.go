package domain

type Question struct {
	id      string
	text    string
	options []Option
	answer  string
}

func (q Question) Id() string {
	return q.id
}

func (q Question) Text() string {
	return q.text
}

func (q Question) Options() []Option {
	return q.options
}

func (q Question) Answer() string {
	return q.answer
}

func NewQuestion(id string, text string, options []Option, answer string) Question {
	return Question{id: id, text: text, options: options, answer: answer}
}
