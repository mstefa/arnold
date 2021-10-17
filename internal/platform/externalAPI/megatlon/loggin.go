package megatlon

import (
	"arnold/internal/gym"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type ExternalSessionClient struct {
	url    string
	method string
}

func NewExternalSessionClient() *ExternalSessionClient {

	return &ExternalSessionClient{
		url:    "https://classes.megatlon.com.ar/api/service/class/club/category/list",
		method: "POST",
	}
}

func (c *ExternalSessionClient) getToken(ctx context.Context, user User) (gym.ExternalSession, error) {

	payload := strings.NewReader("username=mstefanutti24@gmail.com&password=Vsq6Q#ui3xp8pWg&grant_type=password&app_version=1.0")

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
	return gym.ExternalSession{}, err
}
