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

// SignUpWithEmailPassword.
// ref: https://firebase.google.com/docs/reference/rest/auth#section-create-email-password
func (fb *Firebase) SignUpWithEmailPassword(
	ctx context.Context,
	req SignUpWithEmailPasswordRequest,
) (*SignUpWithEmailPasswordResponse, error) {
	url := fmt.Sprintf("%s/v1/accounts:signUp?key=%s", fb.endpoint, fb.apikey)

	return post[SignUpWithEmailPasswordRequest, SignUpWithEmailPasswordResponse](ctx, url, req)
}

// SignInWithEmailPassword.
// ref: https://firebase.google.com/docs/reference/rest/auth#section-sign-in-email-password
func (fb *Firebase) SignInWithEmailPassword(
	ctx context.Context,
	req SignInWithEmailPasswordRequest,
) (*SignInWithEmailPasswordResponse, error) {
	url := fmt.Sprintf("%s/v1/accounts:signInWithPassword?key=%s", fb.endpoint, fb.apikey)

	return post[SignInWithEmailPasswordRequest, SignInWithEmailPasswordResponse](ctx, url, req)
}

func post[Req any, Res any](
	ctx context.Context,
	url string,
	req Req,
) (*Res, error) {
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
		var er ErrorResponse

		if err := json.NewDecoder(res.Body).Decode(&er); err != nil {
			return nil, err
		}

		return nil, &er
	}

	var response Res

	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}

	return &response, nil
}
