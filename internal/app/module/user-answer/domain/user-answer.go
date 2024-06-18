package domain

type UserAnswers struct {
	UserId  string
	Answers []*UserAnswer
}

type UserAnswer struct {
	answerId           string
	userSelectedAnswer string
}

func NewUserAnswer(answerId, userSelectedAnswer string) *UserAnswer {
	return &UserAnswer{answerId: answerId, userSelectedAnswer: userSelectedAnswer}
}

func NewUserAnswersFromRaw(raw map[string]string) []*UserAnswer {
	var answers []*UserAnswer
	for k, v := range raw {
		answers = append(answers, NewUserAnswer(k, v))
	}
	return answers
}

func (u UserAnswer) AnswerId() string {
	return u.answerId
}

func (u UserAnswers) Score() int {
	correctAnswers := 0
	for _, answer := range u.Answers {
		if answer.userSelectedAnswer == answer.AnswerId() {
			correctAnswers++
		}
	}

	return correctAnswers
}
