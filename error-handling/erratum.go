package erratum

import "errors"

// Use opens a resource, calls Frob(input) on the result resource and then closes that resource
func Use(o ResourceOpener, input string) error {
	r, err := tryOpenResource(o)
	if err != nil {
		return err
	}
	defer r.Close()

	return tryFrob(r, input)
}

func tryFrob(r Resource, input string) (result error) {
	defer func() {
		if err := recover(); err != nil {
			switch err.(type) {
			case FrobError:
				fErr := err.(FrobError)
				r.Defrob(fErr.defrobTag)
				result = fErr.inner
			case error:
				result = err.(error)
			default:
				result = errors.New("Unknown error while calling Frob()")
			}
		}
	}()

	r.Frob(input)

	return
}

// tryOpenResource keeps trying to open resource, ignoring TransientError
func tryOpenResource(o ResourceOpener) (Resource, error) {
	for {
		r, err := o()
		if err != nil {
			switch err.(type) {
			case TransientError:
				continue
			default:
				return nil, err
			}
		}
		return r, nil
	}
}
