package main

import (
	"fmt"
	"os"

	"github.com/hellobchain/convertapi-go/pkg"
	"github.com/hellobchain/convertapi-go/pkg/config"
	"github.com/hellobchain/convertapi-go/pkg/param"
)

func main() {
	config.Default = config.NewDefault(os.Getenv("CONVERTAPI_SECRET")) // Get your secret at https://www.convertapi.com/a

	fmt.Println("Converting remote PPTX to PDF")
	pptxRes := convertapi.ConvDef("pptx", "pdf",
		param.NewString("file", "https://cdn.convertapi.com/public/files/demo.pptx"))

	if files, errs := pptxRes.ToPath("/tmp/converted.pdf"); errs == nil {
		fmt.Println("PDF file saved to: ", files[0].Name())
	} else {
		fmt.Println(errs)
	}
}
