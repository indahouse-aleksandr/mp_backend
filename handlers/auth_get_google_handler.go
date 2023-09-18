package handlers

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type GoogleConfig struct {
	Web GoogleConfigNested `json:"web"`
}

type GoogleConfigNested struct {
	ClientId                string   `json:"client_id"`
	ProjectId               string   `json:"project_id"`
	AuthUri                 string   `json:"auth_uri"`
	TokenUri                string   `json:"token_uri"`
	AuthProviderX509CertUrl string   `json:"auth_provider_x509_cert_url"`
	ClientSecret            string   `json:"client_secret"`
	RedirectUris            []string `json:"redirect_uris"`
}

func AuthGetGoogle(w http.ResponseWriter, r *http.Request) {
	oauthState := makeStateOauthCookie(w)
	u := MakeOauth2Config().AuthCodeURL(oauthState)
	http.Redirect(w, r, u, http.StatusTemporaryRedirect)
}

func makeStateOauthCookie(w http.ResponseWriter) string {
	var expiration = time.Now().Add(20 * time.Minute)
	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	cookie := http.Cookie{
		Name:    "oauthstate",
		Value:   state,
		Expires: expiration,
	}
	http.SetCookie(w, &cookie)

	return state
}

func MakeOauth2Config() *oauth2.Config {
	config := readGoogleConfig()
	return &oauth2.Config{
		ClientID:     config.ClientId,
		ClientSecret: config.ClientSecret,
		RedirectURL:  config.RedirectUris[0],
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}
}

func readGoogleConfig() GoogleConfigNested {
	configStr := `{
		"web":{
			"client_id":"",
			"project_id":"",
			"auth_uri":"",
			"token_uri":"",
			"auth_provider_x509_cert_url":"",
			"client_secret":"",
			"redirect_uris":[""]
			}
	}`

	config := GoogleConfig{}
	json.Unmarshal([]byte(configStr), &config)

	return config.Web
}
