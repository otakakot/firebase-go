package firebase

import "errors"

// FirebaseErrorCode is the error code returned by the Firebase API.
type FirebaseErrorCode string

const (
	EmailNotFound            FirebaseErrorCode = "EMAIL_NOT_FOUND"             // There is no user record corresponding to this identifier. The user may have been deleted.
	InvalidPassword          FirebaseErrorCode = "INVALID_PASSWORD"            // The password is invalid or the user does not have a password.
	UserDisabled             FirebaseErrorCode = "USER_DISABLED"               // The user account has been disabled by an administrator.
	EmailExists              FirebaseErrorCode = "EMAIL_EXISTS"                // The email address is already in use by another account.
	CreationNotAlloed        FirebaseErrorCode = "OPERATION_NOT_ALLOWED"       // Password sign-in is disabled for this project.
	TooManyAttemptsTrayLater FirebaseErrorCode = "TOO_MANY_ATTEMPTS_TRY_LATER" // We have blocked all requests from this device due to unusual activity. Try again later.
)

var (
	ErrEmailNotFound            = errors.New(string(EmailNotFound))            // There is no user record corresponding to this identifier. The user may have been deleted.
	ErrInvalidPassword          = errors.New(string(InvalidPassword))          // The password is invalid or the user does not have a password.
	ErrUserDisabled             = errors.New(string(UserDisabled))             // The user account has been disabled by an administrator.
	ErrEmailExists              = errors.New(string(EmailExists))              // The email address is already in use by another account.
	ErrCreationNotAlloed        = errors.New(string(CreationNotAlloed))        // Password sign-in is disabled for this project.
	ErrTooManyAttemptsTrayLater = errors.New(string(TooManyAttemptsTrayLater)) // We have blocked all requests from this device due to unusual activity. Try again later.
)

// ErrorResponse is the response body for the SignInWithEmailPassword method when an error occurs.
// ref: https://firebase.google.com/docs/reference/rest/auth#section-error-response
type ErrorResponse struct {
	Err struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Errors  []struct {
			Message string `json:"message"`
			Domain  string `json:"domain"`
			Reason  string `json:"reason"`
		} `json:"errors"`
	} `json:"error"`
}

// FirebaseError returns the error code as an error.
func (er *ErrorResponse) FirebaseError() error {
	switch er.Err.Message {
	case "EMAIL_NOT_FOUND":
		return ErrEmailNotFound
	case "INVALID_PASSWORD":
		return ErrInvalidPassword
	case "USER_DISABLED":
		return ErrUserDisabled
	case "EMAIL_EXISTS":
		return ErrEmailExists
	case "OPERATION_NOT_ALLOWED":
		return ErrCreationNotAlloed
	case "TOO_MANY_ATTEMPTS_TRY_LATER":
		return ErrTooManyAttemptsTrayLater
	default:
		return errors.New(er.Err.Message)
	}
}

// Error returns the error message.
func (er *ErrorResponse) Error() string {
	switch er.Err.Message {
	case "EMAIL_NOT_FOUND":
		return ErrEmailNotFound.Error()
	case "INVALID_PASSWORD":
		return ErrInvalidPassword.Error()
	case "USER_DISABLED":
		return ErrUserDisabled.Error()
	case "EMAIL_EXISTS":
		return ErrEmailExists.Error()
	case "OPERATION_NOT_ALLOWED":
		return ErrCreationNotAlloed.Error()
	case "TOO_MANY_ATTEMPTS_TRY_LATER":
		return ErrTooManyAttemptsTrayLater.Error()
	default:
		return er.Err.Message
	}
}
