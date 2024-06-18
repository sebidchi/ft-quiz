package test

type InvalidQuery struct {
}

func (i InvalidQuery) Id() string {
	return "InvalidQuery"
}

type InvalidCommand struct {
}

func (i InvalidCommand) Id() string {
	return "invalid-command"
}
