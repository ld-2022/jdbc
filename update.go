package main

type UpdateOption func(*UpdateOptions)
type UpdateOptions struct {
	Sql  string
	Args []any
}

func loadUpOp(option ...UpdateOption) *UpdateOptions {
	options := new(UpdateOptions)
	for _, e := range option {
		e(options)
	}
	return options
}
