package vault

import (
	"encoding/json"
	"net/http"
	"workspace-service/core/config"

	"golang.org/x/oauth2"
)

type VaultTokenSource struct {
	VaultToken                    string
	VaultTokenHeader              string
	VaultAdminManagementClientUri string
}

type VaultQueryResponse struct {
	Data oauth2.Token
}

func (vts *VaultTokenSource) Token() (*oauth2.Token, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", vts.VaultAdminManagementClientUri, nil)
	req.Header.Add(vts.VaultTokenHeader, vts.VaultToken)
	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var token VaultQueryResponse
	err = json.NewDecoder(res.Body).Decode(&token)
	if err != nil {
		return nil, err
	}
	return &token.Data, nil
}

func GetTokenSource(config config.VaultConfig) *VaultTokenSource {
	return &VaultTokenSource{
		VaultToken:                    config.Token,
		VaultTokenHeader:              config.TokenHeader,
		VaultAdminManagementClientUri: config.AdminManagementClientUri,
	}
}
