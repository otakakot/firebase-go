package firebase

import "errors"

// SignInWithEmailPasswordRequest is the request body for the SignInWithEmailPassword method.
type SignInWithEmailPasswordRequest struct {
	Email             string `json:"email"`             // The email the user is signing in with.
	Password          string `json:"password"`          // The password for the account.
	ReturnSerureToken bool   `json:"returnSecureToken"` // Whether or not to return an ID and refresh token. Should always be true.
}

// SignInWithEmailPasswordResponse is the response body for the SignInWithEmailPassword method.
type SignInWithEmailPasswordResponse struct {
	IDToken      string `json:"idToken"`      // A Firebase Auth ID token for the authenticated user.
	Email        string `json:"email"`        // The email for the authenticated user.
	RefreshToken string `json:"refreshToken"` // A Firebase Auth refresh token for the authenticated user.
	ExpiresIn    string `json:"exiresIn"`     // The number of seconds in which the ID token expires.
	LocalID      string `json:"localId"`      // The uid of the authenticated user.
	Registered   bool   `json:"registered"`   // Whether the email is for an existing account.
}

var (
	ErrEmailNotFound   = errors.New("EMAIL_NOT_FOUND")  // There is no user record corresponding to this identifier. The user may have been deleted.
	ErrInvalidPassword = errors.New("INVALID_PASSWORD") // The password is invalid or the user does not have a password.
	ErrUserDisabled    = errors.New("USER_DISABLED")    // The user account has been disabled by an administrator.
)

// errorResponse is the response body for the SignInWithEmailPassword method when an error occurs.
type errorResponse struct {
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

func (e errorResponse) Error() string {
	switch e.Err.Message {
	case "EMAIL_NOT_FOUND":
		return ErrEmailNotFound.Error()
	case "INVALID_PASSWORD":
		return ErrInvalidPassword.Error()
	case "USER_DISABLED":
		return ErrUserDisabled.Error()
	default:
		return e.Err.Message
	}
}
