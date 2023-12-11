package network

import (
	"fmt"
	"net/http"
)

const (
	authorization_code = "authorization_code"

	authorization         = "Authorization"
	content_type          = "Content-Type"
	refresh_token         = "refresh_token"
	delete                = "delete"
	x_www_form_urlencoded = "application/x-www-form-urlencoded"
)

func LoginAuthentication(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://prd.kr-ccapi.hyundai.com/api/v1/user/oauth2/authorize?client_id=81299d4d-0b91-425d-80d9-32c4e78d68a2&redirect_uri=http://localhost:8080&response_type=code&state=test", http.StatusFound)
}

func IssueToken(w http.ResponseWriter, r *http.Request) {
	redirectUrl := fmt.Sprintf(
		"https://prd.kr-ccapi.hyundai.com/api/v1/user/oauth2/token?grant_type=%s&code=%s&redirect_uri=%s",
		authorization_code,
		r.URL.Query().Get("code"),
		"http://localhost:8080/token",
	)
	fmt.Println(redirectUrl)

	client := &http.Client{}
	req, _ := http.NewRequest("POST", redirectUrl, nil)
	req.Header.Set(authorization, "ODEyOTlkNGQtMGI5MS00MjVkLTgwZDktMzJjNGU3OGQ2OGEyOnR2RElhWTVaZFAxMkZ4ZVNZNU9yUzFsb2ZNMlFRSWVOYVkyU2RqREgxMGlGMkl2Wg==")
	req.Header.Set(content_type, x_www_form_urlencoded)
	response, _ := client.Do(req)

	fmt.Println(response)
}
