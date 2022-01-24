package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
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
	urlString := flag.Arg(0)

	if len(urlString) == 0 {
		help()
		return 0
	}

	parsedUrl, err := url.Parse(urlString)
	if err != nil {
		println(err.Error())
		return 1
	}
	if parsedUrl.Scheme == "" {
		parsedUrl.Scheme = "http"
	}

	req, err := http.NewRequest(method, parsedUrl.String(), nil)
	if err != nil {
		println(err.Error())
		return 1
	}

	if verbose {
		println("method:", req.Method)
		println("url:", req.URL.String())
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		println(err.Error())
		return 1
	}
	defer res.Body.Close()

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
