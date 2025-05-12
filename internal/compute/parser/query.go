package parser

type Query struct {
	command string
	args    []string
}

func (q *Query) Command() string {
	return q.command
}

func (q *Query) Args() []string {
	return q.args
}
