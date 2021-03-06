package vault

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	log "github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jwt"
	"google.golang.org/api/iam/v1"
)

// GCPBackendConfig parmaters for GCP backend login through Vault
type GCPBackendConfig struct {
	Project        string
	CredsPath      string
	ServiceAccount string
}

// GetServiceAccountCreds read the service account json
func GetServiceAccountCreds(cfg *GCPBackendConfig) (*jwt.Config, error) {
	log.Info("Getting service account credential file...")
	jsonBytes, err := ioutil.ReadFile(cfg.CredsPath)
	if err != nil {
		log.Fatalf("error reading credentials file %v", err)
	}
	config, err := google.JWTConfigFromJSON(jsonBytes, iam.CloudPlatformScope)
	if err != nil {
		return nil, fmt.Errorf("error getting JWT config from JSON %v", err)
	}
	return config, nil
}

func generateSignedJWTWithIAM(iamClient *iam.Service, cfg *GCPBackendConfig, role string) (*iam.SignJwtResponse, error) {
	resourceName := fmt.Sprintf("projects/%s/serviceAccounts/%s", cfg.Project, cfg.ServiceAccount)
	log.Debugf("Generating signed JWT with IAM for resource %s", resourceName)
	jwtPayload := map[string]interface{}{
		"sub": cfg.ServiceAccount,
		"aud": fmt.Sprintf("vault/%s", role),
		"exp": time.Now().Add(time.Minute * 10).Unix(),
	}

	payloadBytes, err := json.Marshal(jwtPayload)
	if err != nil {
		return nil, fmt.Errorf("error decoding JSON payload from vault gcp login %v", err)
	}

	signJwtReq := &iam.SignJwtRequest{
		Payload: string(payloadBytes),
	}

	resp, err := iamClient.Projects.ServiceAccounts.SignJwt(resourceName, signJwtReq).Do()
	if err != nil {
		return nil, fmt.Errorf("error siging JWT using GCP IAM %v", err)
	}
	return resp, nil
}

// GCPBackendLogin Authenticate to Vault via GCP Backend
func GCPBackendLogin(client *Client, gcpBackendConfig *GCPBackendConfig, vaultConfig *Config) (string, error) {
	var logger *log.Entry

	config, err := GetServiceAccountCreds(gcpBackendConfig)
	if err != nil {
		return "", err
	}
	gcpBackendConfig.ServiceAccount = config.Email
	httpClient := config.Client(oauth2.NoContext)
	iamClient, err := iam.New(httpClient)
	if err != nil {
		return "", err
	}

	resp, err := generateSignedJWTWithIAM(iamClient, gcpBackendConfig, vaultConfig.Role)
	if err != nil {
		return "", err
	}
	// Send signed JWT in login request to Vault.
	params := map[string]interface{}{
		"role": vaultConfig.Role,
		"jwt":  resp.SignedJwt,
	}
	logger = log.WithFields(log.Fields{
		"project":        gcpBackendConfig.Project,
		"serviceAccount": gcpBackendConfig.ServiceAccount,
	})
	logger.Infof("Login into Vault GCP backend using the role %s", vaultConfig.Role)
	secretData, err := client.Logical.Write("auth/gcp/login", params)
	if err != nil {
		return "", fmt.Errorf("failed login to Vault using GCP backend %v", err)
	}
	clientToken := &secretData.Auth.ClientToken
	return *clientToken, nil
}
