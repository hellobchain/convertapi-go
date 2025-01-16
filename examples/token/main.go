package main

import (
	"fmt"
	"os"

	convertapi "github.com/hellobchain/convertapi-go/pkg"
	"github.com/hellobchain/convertapi-go/pkg/config"
)

func main() {
	config.Default = config.NewDefault(os.Getenv("CONVERTAPI_TOKEN")) // Get your secret at https://www.convertapi.com/a/access-tokens

	if file, errs := convertapi.ConvertPath("assets/test.docx", "/tmp/result.pdf"); errs == nil {
		fmt.Println("PDF file saved to: ", file.Name())
	} else {
		fmt.Println(errs)
	}
}
