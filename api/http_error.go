package franky

type HttpError struct {
	StatusCode int
	Err        error
}
