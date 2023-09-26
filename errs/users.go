package errs

const UserEmaiExistMsg = "user email already exists"

type ErrorUserEmaiExist struct {
	fieldName string
	err       string
}

func (e *ErrorUserEmaiExist) Field() string {
	return e.fieldName
}

func (e *ErrorUserEmaiExist) Error() string {
	return e.err
}

func NewUserEmaiExist(fieldName string, err string) ErrorUserEmaiExist {
	return ErrorUserEmaiExist{
		fieldName: fieldName,
		err:       err,
	}
}
