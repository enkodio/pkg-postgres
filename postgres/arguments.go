package postgres

type Arguments struct {
	p    Placeholder
	args []interface{}
}

func NewArguments(values ...interface{}) Arguments {
	a := Arguments{
		args: make([]interface{}, 0),
		p:    Placeholder(0),
	}
	if len(values) > 0 {
		a.args = append(a.args, values...)
	}
	a.p = Placeholder(len(values))
	return a
}

func (a *Arguments) Add(v interface{}) string {
	a.args = append(a.args, v)
	return a.p.Next()
}

func (a *Arguments) Get() []interface{} {
	return a.args
}
