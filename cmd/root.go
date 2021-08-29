// Copyright © 2021 Juanpis
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/jochasinga/requests"
	"github.com/spf13/cobra"
)

// Merges all the strings that start from given argument position into one variable.
func mergeStrings(args []string, arg_location int) string {
	var msg string
	for i := arg_location; i < len(args); i++ {
		msg = msg + " " + args[i]
	}
	msg = strings.TrimSpace(msg)
	return msg
}

// Does a HTTP request method to Discord's webhook with provided URL and JSON map.
// Automatically marshalls the JSON map.
//
// Supported HTTP Methods: POST, PATCH
func requestHTTP(HttpMethod string, URL string, JsonMap map[string]string) {
	jsonValue, _ := json.Marshal(JsonMap)

	switch HttpMethod {
	case "POST":
		resp, err := requests.Post(URL, "application/json", bytes.NewBuffer(jsonValue))
		_, _ = resp, err
	case "PATCH":
		resp, err := requests.Patch(URL, "application/json", bytes.NewBuffer(jsonValue))
		_, _ = resp, err
	}
}

// Checks if provided URL matches the Discord URL webhook API calling one.
func isTokenValid(url string) bool {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("ERROR: '%s' is not a valid webhook URL.", url)
			os.Exit(0)
		}
	}()

	if url[0:33] == "https://discord.com/api/webhooks/" {
		url_r, err := requests.Get(url)
		url_code := url_r.StatusCode
		if err != nil || url_code == 401 {
			return false
		} else {
			return true
		}
	} else {
		return false
	}
}

// Checks if message argument doesn't pass 2000 characters.
func isMsgMax(msg string) bool {
	msgLen := len(msg)
	msgLimit := 2000 // you never know if discord may change their
	// limit in the near future /shrug

	if msgLen < msgLimit {
		return false
	} else {
		// msgToShort := msgLen - msgLimit
		// fmt.Printf("Your message's length (%d) surpasses Discord's limit (%d)."+
		// 	"Please make it %d characters shorter and try again.",
		// 	msgLen, msgLimit, msgToShort)
		return true
	}
}

// Panics when an error ocurrs. Only use if no conditionals are being used or if it's a HTTPS request.
func ManageError(err error) {
	// just to calm down with the syntax
	if err != nil {
		fmt.Println("An unexpected error ocurred. Please try again.")
		fmt.Println("For more information:")
		panic(err)
	}
}

// Cobra stuff

var rootCmd = &cobra.Command{
	Use:  "dishook [url] [message]\n  dishook [url] [flags]",
	Args: cobra.MinimumNArgs(2),

	PreRunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			cmd.Help()
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		execute(args)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
