package pkg

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlayIntegrityVerdictClient_DecryptVerdict_Success(t *testing.T) {

	testToken := "this is a test token"
	testVerdictResponse := &integrityTokenDecryptionResponse{
		TokenPayloadExternal: &PlayIntegrityVerdict{
			RequestDetails: &RequestDetails{RequestPackageName: "test"},
		},
	}

	is := assert.New(t)
	must := assert.New(t)

	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		is.Equal(http.MethodPost, r.Method)
		body, err := ioutil.ReadAll(r.Body)
		must.NoError(err)

		tokenRequest := &integrityTokenDecryptionRequest{}
		err = json.Unmarshal(body, tokenRequest)
		must.NoError(err)

		is.Equal(testToken, tokenRequest.IntegrityToken)
		responseBytes, err := json.Marshal(testVerdictResponse)
		must.NoError(err)

		_, err = w.Write(responseBytes)
		must.NoError(err)
	}))
	defer svr.Close()

	verdictClient := &PlayIntegrityVerdictClient{
		client: http.DefaultClient,
		url:    svr.URL,
	}

	verdict, err := verdictClient.DecryptVerdict(context.Background(), testToken)
	is.NoError(err)
	is.Equal(testVerdictResponse.TokenPayloadExternal, verdict)
}

func TestPlayIntegrityVerdictClient_DecryptVerdict_ErrorResponse(t *testing.T) {
	is := assert.New(t)

	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer svr.Close()

	verdictClient := &PlayIntegrityVerdictClient{
		client: http.DefaultClient,
		url:    svr.URL,
	}

	_, err := verdictClient.DecryptVerdict(context.Background(), "test")
	is.EqualError(err, "token decryption request failed: receive status 500 Internal Server Error")
}
