package megatlon

import (
	"arnold/internal/gym"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type ExternalSessionClient struct {
	url    string
	method string
}

type Response struct {
	AccessToken  string `json:"acces_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	Scope        string `json:"scope"`
}

func NewExternalSessionClient() *ExternalSessionClient {

	return &ExternalSessionClient{
		url:    "https://users.megatlon.com.ar/oauth/token",
		method: "POST",
	}
}

func (c *ExternalSessionClient) GetToken(user gym.User) (gym.ExternalSession, error) {

	payload := strings.NewReader("username=mstefanutti24%40gmail.com&password=Vsq6Q%23ui3xp8pWg&grant_type=password&app_version=1.0")

	client := &http.Client{}
	req, err := http.NewRequest(c.method, c.url, payload)

	if err != nil {
		fmt.Println(err)
		return gym.ExternalSession{}, err
	}
	req.Header.Add("authorization", "Basic bWVnYXRsb246dXNlcg==") //get from btoa('megatlon:user')
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return gym.ExternalSession{}, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return gym.ExternalSession{}, err
	}
	fmt.Println(string(body))

	var responseObject Response
	json.Unmarshal(body, &responseObject)
	fmt.Printf("API Response as struct %+v\n", responseObject)

	r, err := gym.NewExternalSession(user.ID.String(), user.ID, responseObject.AccessToken, responseObject.RefreshToken, responseObject.Scope, responseObject.TokenType)

	return r, err
}
