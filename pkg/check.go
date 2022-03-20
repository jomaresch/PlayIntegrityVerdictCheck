package pkg

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/oauth2/google"
)

const (
	integrityScope = "https://www.googleapis.com/auth/playintegrity"

	integrityAPIFormat = "https://playintegrity.googleapis.com/v1/%s:decodeIntegrityToken"
)

type PlayIntegrityVerdictClient struct {
	client *http.Client
	url    string
}

func NewPlayIntegrityVerdictClient(ctx context.Context, appPackageName string) (*PlayIntegrityVerdictClient, error) {
	client, err := google.DefaultClient(ctx, integrityScope)
	if err != nil {
		return nil, err
	}
	return &PlayIntegrityVerdictClient{client: client, url: fmt.Sprintf(integrityAPIFormat, appPackageName)}, nil
}

func (p *PlayIntegrityVerdictClient) DecryptVerdict(ctx context.Context, verdictToken string) (*PlayIntegrityVerdict, error) {
	tokenDecryptionRequest := integrityTokenDecryptionRequest{IntegrityToken: verdictToken}
	requestBody, err := json.Marshal(tokenDecryptionRequest)
	if err != nil {
		return nil, fmt.Errorf("failed to marshall request token: %w", err)
	}
	request, err := http.NewRequest(http.MethodPost, p.url, bytes.NewReader(requestBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create http request: %w", err)
	}
	request.WithContext(ctx)
	response, err := p.client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("failed to do token decryption request: %w", err)
	}
	defer func() {
		err := response.Body.Close()
		if err != nil {
			log.Printf("failed to close body: %v", err)
		}
	}()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("token decryption request failed: receive status %s", response.Status)
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read body: %w", err)
	}
	verdictResponse := &integrityTokenDecryptionResponse{}

	err = json.Unmarshal(body, verdictResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshall verdict: %w", err)
	}
	return verdictResponse.TokenPayloadExternal, nil
}
