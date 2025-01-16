package convertapi

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/hellobchain/convertapi-go/pkg/config"
	"github.com/hellobchain/convertapi-go/pkg/lib"
)

type User struct {
	Active      bool
	ApiKey      int
	Email       string
	FullName    string
	SecondsLeft int
	Secret      string
	Status      string
}

func UserInfo(conf *config.Config) (user *User, err error) {
	if conf == nil {
		conf = config.Default
	}
	query := url.Values{}
	path := fmt.Sprintf("/user?%s", query.Encode())
	pathURL, err := url.Parse(path)
	if err != nil {
		return
	}

	userURL := conf.BaseURL.ResolveReference(pathURL)
	resp, err := lib.RespExtractErr(conf.HttpClient.Get(userURL.String()))
	if err == nil {
		defer resp.Body.Close()
		user = &User{}
		err = json.NewDecoder(resp.Body).Decode(user)
	}

	return
}
