package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/tomon9086/tapi/src/request"
)

const CMD_NAME = "tapi"

func help() {
	os.Stdout.Write([]byte(fmt.Sprintf("Usage of %s:\n", CMD_NAME)))
	flag.PrintDefaults()
}

func m() int {
	flag.Usage = help

	pMethod := flag.String("m", "get", "request method")
	pVerbose := flag.Bool("v", false, "verbose")
	flag.Parse()

	method := strings.ToUpper(*pMethod)
	verbose := *pVerbose
	url := flag.Arg(0)

	req, res, err := request.Request(url, request.RequestOption{
		Method: method,
	})
	if err != nil {
		if err == request.ErrRequestEmptyUrl {
			help()
		} else {
			println(err.Error())
		}
		return 1
	}
	defer res.Body.Close()

	if verbose {
		println("method:", req.Method)
		println("url:", req.URL.String())
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		println(err.Error())
		return 1
	}
	_, err = os.Stdout.Write(body)
	if err != nil {
		println(err.Error())
		return 1
	}

	return 0
}

func main() {
	os.Exit(m())
}
