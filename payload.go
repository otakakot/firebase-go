package firebase

type SignUpWithEmailPasswordRequest struct {
	Email             string `json:"email"`             // The email for the user to create.
	Password          string `json:"password"`          // The password for the user to create.
	ReturnSerureToken bool   `json:"returnSecureToken"` // Whether or not to return an ID and refresh token. Should always be true.
}

type SignUpWithEmailPasswordResponse struct {
	IDToken      string `json:"idToken"`      // A Firebase Auth ID token for the newly created user.
	Email        string `json:"email"`        // The email for the newly created user.
	RefreshToken string `json:"refreshToken"` // A Firebase Auth refresh token for the newly created user.
	ExpiresIn    string `json:"expiresIn"`    // The number of seconds in which the ID token expires.
	LocalID      string `json:"localId"`      // The uid of the newly created user.
}

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
	ExpiresIn    string `json:"expiresIn"`    // The number of seconds in which the ID token expires.
	LocalID      string `json:"localId"`      // The uid of the authenticated user.
	Registered   bool   `json:"registered"`   // Whether the email is for an existing account.
}
