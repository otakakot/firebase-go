package firebase

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

const endpoint = "https://identitytoolkit.googleapis.com"

type Options func(*Firebase)

func Endpoint(endpoint string) Options {
	return func(f *Firebase) {
		f.endpoint = endpoint
	}
}

type Firebase struct {
	endpoint string
	apikey   string
}

func New(apikey string, options ...Options) *Firebase {
	fb := &Firebase{
		apikey:   apikey,
		endpoint: endpoint,
	}

	for _, option := range options {
		option(fb)
	}

	return fb
}

// SignInWithEmailPassword.
// https://firebase.google.com/docs/reference/rest/auth#section-sign-in-email-password
func (fb *Firebase) SignInWithEmailPassword(
	ctx context.Context,
	req SignInWithEmailPasswordRequest,
) (*SignInWithEmailPasswordResponse, error) {
	url := fmt.Sprintf("%s/v1/accounts:signInWithPassword?key=%s", fb.endpoint, fb.apikey)

	body := new(bytes.Buffer)

	if err := json.NewEncoder(body).Encode(req); err != nil {
		return nil, err
	}

	request, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(request.WithContext(ctx))
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		var er errorResponse

		if err := json.NewDecoder(res.Body).Decode(&er); err != nil {
			return nil, err
		}

		return nil, er
	}

	var response SignInWithEmailPasswordResponse

	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}

	return &response, nil
}
