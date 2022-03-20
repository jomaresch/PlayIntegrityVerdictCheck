# Play Integrity Verdict Check

This is a small library which calls the Google API to decrypt and verify a Play Integrity Token.

### Play Integrity API 
https://developer.android.com/google/play/integrity/overview

# Install

```
go get github.com/jomaresch/PlayIntegrityVerdictCheck
```

# Usage

The library uses the default Google Http Client, so you have multiple ways for authentication.
https://cloud.google.com/docs/authentication/production

```go

// (optional) Set service account path.
pathToServiceAccount := "/app/service-account.json"
os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", pathToServiceAccount)

// Create a new client.
appPackageName := "com.example.myapp"
client, err := NewPlayIntegrityVerdictClient(ctx, appPackageName)
if err != nil {
    panic(err)
}

// Decrypt your token. 
token := "ADD_YOUR_TOKEN"
verdict, err := client.DecryptVerdict(ctx2, token)
if err != nil {
    panic(err)
}
```
