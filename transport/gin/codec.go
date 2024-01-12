package gin

type EncodeResponse (func(Context, interface{}) error)

type EncodeError (func(Context, error) error)

func DefaultEncodeResponse(ctx Context, i interface{}) error {
	return nil
}

func DefaultEncodeError(ctx Context, err error) error {
	return nil
}
