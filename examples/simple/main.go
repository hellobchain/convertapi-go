package main

import (
	"fmt"
	"os"

	"github.com/hellobchain/convertapi-go/pkg"
	"github.com/hellobchain/convertapi-go/pkg/config"
)

func main() {
	config.Default = config.NewDefault(os.Getenv("CONVERTAPI_SECRET")) // Get your secret at https://www.convertapi.com/a
	if file, errs := convertapi.ConvertPath("assets/test.docx", "/tmp/result.pdf"); errs == nil {
		fmt.Println("PDF file saved to: ", file.Name())
	} else {
		fmt.Println(errs)
	}
}
