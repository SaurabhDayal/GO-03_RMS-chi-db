package errorHandling

import "06_RMS-chi-db/models"

func UserAlreadyPresent() error {
	return &models.CustomClientError{
		Message:    "user already present",
		StatusCode: 400,
	}
}

func IncorrectEmail() error {
	return &models.CustomClientError{
		Message:    "Incorrect email",
		StatusCode: 400,
	}
}

func UserCredentialNotMatch() error {
	return &models.CustomClientError{
		Message:    "Incorrect user email and password",
		StatusCode: 400,
	}
}

func InvalidUser() error {
	return &models.CustomClientError{
		Message:    "Invalid Authorization",
		StatusCode: 400,
	}
}

func UnableToAccessDB() error {
	return &models.CustomClientError{
		Message:    "unable to access into DB",
		StatusCode: 500,
	}
}

func UnableToReadURL() error {
	return &models.CustomClientError{
		Message:    "unable to convert URL path variables into integers ",
		StatusCode: 400,
	}
}

func noContentDB() error {
	return &models.CustomClientError{
		Message:    "no content in DB",
		StatusCode: 204,
	}
}

func UnableToReadJSON() error {
	return &models.CustomClientError{
		Message:    "unable to read or write JSON",
		StatusCode: 400,
	}
}

func UnableToWriteJSON() error {
	return &models.CustomClientError{
		Message:    "unable to read or write JSON",
		StatusCode: 500,
	}
}

func UnableToGenerateToken() error {
	return &models.CustomClientError{
		Message:    "unable to generate token",
		StatusCode: 500,
	}
}
