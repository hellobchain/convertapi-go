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

	fmt.Println("Creating PDF with the first and the last pages")
	splitRes := convertapi.ConvDef("pdf", "split",
		param.NewPath("file", "assets/test.pdf", nil),
	)

	mergeRes := convertapi.ConvDef("pdf", "merge",
		param.NewResultIdx("files", splitRes, 0, nil),
		param.NewResultIdx("files", splitRes, -1, nil),
	)

	if files, errs := mergeRes.ToPath("/tmp"); errs == nil {
		fmt.Println("PDF file saved to: ", files[0].Name())
	} else {
		fmt.Println(errs)
	}
}
