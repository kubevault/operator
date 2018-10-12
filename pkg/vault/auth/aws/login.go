package aws

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
	vaultapi "github.com/hashicorp/vault/api"
	"github.com/hashicorp/vault/helper/awsutil"
	"github.com/kubevault/operator/pkg/vault/auth/types"
	vaultuitl "github.com/kubevault/operator/pkg/vault/util"
	"github.com/pkg/errors"
	core "k8s.io/api/core/v1"
	appcat "kmodules.xyz/custom-resources/apis/appcatalog/v1alpha1"
)

const iamServerIdHeader = "X-Vault-AWS-IAM-Server-ID"

type auth struct {
	vClient     *vaultapi.Client
	creds       *credentials.Credentials
	headerValue string
	role        string
}

type params struct {
	Role        string `json:"role"`
	HeaderValue string `json:"header_value,omitempty"`
}

// links : https://www.vaultproject.io/docs/auth/aws.html

func New(vApp *appcat.AppBinding, secret *core.Secret) (*auth, error) {
	if vApp.Spec.Parameters == nil {
		return nil, errors.New("parameters are not provided")
	}

	cfg, err := vaultuitl.VaultConfigFromAppBinding(vApp)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create vault config from AppBinding")
	}

	vc, err := vaultapi.NewClient(cfg)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create vault client")
	}

	accessKeyID, ok := secret.Data["access_key_id"]
	if !ok {
		return nil, errors.New("access_key_id is missing")
	}
	secretAccessKey, ok := secret.Data["secret_access_key"]
	if !ok {
		return nil, errors.New("secret_access_key is missing")
	}
	securityToken := secret.Data["security_token"]

	var p params
	err = json.Unmarshal(vApp.Spec.Parameters.Raw, &p)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal parameters")
	}

	creds, err := retrieveCreds(string(accessKeyID), string(secretAccessKey), string(securityToken))
	if err != nil {
		return nil, errors.Wrap(err, "failed to retrieve credentials")
	}

	return &auth{
		vClient:     vc,
		creds:       creds,
		role:        string(p.Role),
		headerValue: string(p.HeaderValue),
	}, nil
}

// Login will log into vault and return client token
func (a *auth) Login() (string, error) {
	req := a.vClient.NewRequest("POST", "/v1/auth/aws/login")

	loginData, err := a.generateLoginData()
	if err != nil {
		return "", errors.Wrap(err, "failed to generate login data")
	}

	payload := loginData
	payload["role"] = a.role
	if err := req.SetJSONBody(payload); err != nil {
		return "", err
	}

	resp, err := a.vClient.RawRequest(req)
	if err != nil {
		return "", err
	}

	var loginResp types.AuthLoginResponse
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&loginResp)
	if err != nil {
		return "", err
	}
	return loginResp.Auth.ClientToken, nil
}

// Generates the necessary data to send to the Vault server for generating a token
// This is useful for other API clients to use
func (a *auth) generateLoginData() (map[string]interface{}, error) {
	loginData := make(map[string]interface{})

	// Use the credentials we've found to construct an STS session
	stsSession, err := session.NewSessionWithOptions(session.Options{
		Config: aws.Config{Credentials: a.creds},
	})
	if err != nil {
		return nil, err
	}

	var params *sts.GetCallerIdentityInput
	svc := sts.New(stsSession)
	stsRequest, _ := svc.GetCallerIdentityRequest(params)

	// Inject the required auth header value, if supplied, and then sign the request including that header
	if a.headerValue != "" {
		stsRequest.HTTPRequest.Header.Add(iamServerIdHeader, a.headerValue)
	}
	stsRequest.Sign()

	// Now extract out the relevant parts of the request
	headersJson, err := json.Marshal(stsRequest.HTTPRequest.Header)
	if err != nil {
		return nil, err
	}
	requestBody, err := ioutil.ReadAll(stsRequest.HTTPRequest.Body)
	if err != nil {
		return nil, err
	}
	loginData["iam_http_request_method"] = stsRequest.HTTPRequest.Method
	loginData["iam_request_url"] = base64.StdEncoding.EncodeToString([]byte(stsRequest.HTTPRequest.URL.String()))
	loginData["iam_request_headers"] = base64.StdEncoding.EncodeToString(headersJson)
	loginData["iam_request_body"] = base64.StdEncoding.EncodeToString(requestBody)

	return loginData, nil
}

func retrieveCreds(accessKey, secretKey, sessionToken string) (*credentials.Credentials, error) {
	credConfig := &awsutil.CredentialsConfig{
		AccessKey:    accessKey,
		SecretKey:    secretKey,
		SessionToken: sessionToken,
	}
	creds, err := credConfig.GenerateCredentialChain()
	if err != nil {
		return nil, err
	}
	if creds == nil {
		return nil, errors.New("could not compile valid credential providers from static config, environment, shared, or instance metadata")
	}

	_, err = creds.Get()
	if err != nil {
		return nil, errors.Wrap(err, "failed to retrieve credentials from credential chain")
	}
	return creds, nil
}