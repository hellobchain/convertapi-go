package main

import (
	"fmt"
	"os"

	"github.com/hellobchain/convertapi-go/pkg"
	"github.com/hellobchain/convertapi-go/pkg/config"
)

func main() {
	config.Default = config.NewDefault(os.Getenv("CONVERTAPI_SECRET")) // Get your secret at https://www.convertapi.com/a

	if user, err := convertapi.UserInfo(nil); err == nil {
		fmt.Println("User information: ")
		fmt.Printf("%+v\n", user)
	} else {
		fmt.Println(err)
	}
}
