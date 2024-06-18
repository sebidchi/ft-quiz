package domain

type Option struct {
	id   string
	text string
}

func (o Option) Id() string {
	return o.id
}

func (o Option) Text() string {
	return o.text
}

func NewOption(id string, text string) *Option {
	return &Option{id: id, text: text}
}
