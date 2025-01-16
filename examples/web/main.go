package main

import (
	"fmt"
	"os"

	convertapi "github.com/hellobchain/convertapi-go/pkg"
	"github.com/hellobchain/convertapi-go/pkg/config"
	"github.com/hellobchain/convertapi-go/pkg/param"
)

func main() {
	config.Default = config.NewDefault(os.Getenv("CONVERTAPI_SECRET")) // Get your secret at https://www.convertapi.com/a

	fmt.Println("Converting remote PPTX to PDF")
	pptxRes := convertapi.ConvDef("web", "pdf",
		param.NewString("url", "https://en.wikipedia.org/wiki/Data_conversion"),
		param.NewString("filename", "web-example"),
	)

	if files, errs := pptxRes.ToPath("/tmp"); errs == nil {
		fmt.Println("PDF file saved to: ", files[0].Name())
	} else {
		fmt.Println(errs)
	}
}
