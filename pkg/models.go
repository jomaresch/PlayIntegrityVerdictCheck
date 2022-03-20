package pkg

type PlayIntegrityVerdict struct {
	RequestDetails  *RequestDetails  `json:"requestDetails"`
	AppIntegrity    *AppIntegrity    `json:"appIntegrity"`
	DeviceIntegrity *DeviceIntegrity `json:"deviceIntegrity"`
	AccountDetails  *AccountDetails  `json:"accountDetails"`
}

type RequestDetails struct {
	RequestPackageName string `json:"requestPackageName"`
	Nonce              string `json:"nonce"`
	TimestampMillis    string `json:"timestampMillis"`
}

type AppIntegrity struct {
	CertificateSha256Digest []string `json:"certificateSha256Digest"`
	AppRecognitionVerdict   string   `json:"appRecognitionVerdict"`
	PackageName             string   `json:"packageName"`
	VersionCode             string   `json:"versionCode"`
}

type DeviceIntegrity struct {
	DeviceRecognitionVerdict []string `json:"deviceRecognitionVerdict"`
}

type AccountDetails struct {
	AppLicensingVerdict string `json:"appLicensingVerdict"`
}

type integrityTokenDecryptionRequest struct {
	IntegrityToken string `json:"integrity_token"`
}

type integrityTokenDecryptionResponse struct {
	TokenPayloadExternal *PlayIntegrityVerdict `json:"tokenPayloadExternal"`
}
