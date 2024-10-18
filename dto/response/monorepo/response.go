package response

type OutputFormat struct {
	Success bool
	Message string
	Data    interface{}
	Errors  []struct {
		ErrorFormat
	}
	Code string
}

type ErrorFormat struct {
	Field string
	Error string
}
