package helpers

func Validate(body []byte) bool {
	bodyToValidate := string(body)

	if len(bodyToValidate) != 0 {
		return false
	}

	return true
}
