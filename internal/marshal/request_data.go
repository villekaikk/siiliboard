package marshal

type RequestTemplate interface {
	Validate() error
}
