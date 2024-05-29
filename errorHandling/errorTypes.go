package errorHandling

import "03_RMS/models"

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

func NoContentDB() error {
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

func UnableToBeginTransaction() error {
	return &models.CustomClientError{
		Message:    "unable to begin transaction",
		StatusCode: 500,
	}
}

func UnableToRollbackTransaction() error {
	return &models.CustomClientError{
		Message:    "unable to rollback transaction",
		StatusCode: 500,
	}
}
func UnableToCommitTransaction() error {
	return &models.CustomClientError{
		Message:    "unable to commit transaction",
		StatusCode: 500,
	}
}
