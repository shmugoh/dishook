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

var rootCmd = &cobra.Command{
	Use:   " dishook [url] [message]\nor dishook [url]",
	Short: "CLI webhook runner for Discord.",
	Args:  cobra.ArbitraryArgs,
	// Args:  cobra.MinimumNArgs(2),
	// Long: maybe i won't use it, but i'll leave it here just in case
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		execute(args)
	},
}

// functions time

func getContent(args []string, arg_location int) string {
	var msg string
	for i := arg_location; i < len(args); i++ {
		msg = msg + " " + args[i]
	}
	msg = strings.TrimSpace(msg)
	return msg
}

func sendMsg(url string, content string) {
	values := map[string]string{"content": content}
	jsonValue, _ := json.Marshal(values)
	// Turns content string into JSON

	resp, err := requests.Post(url, "application/json", bytes.NewBuffer(jsonValue))
	fmt.Print(resp)
	manageError(err)
	// Sends request to webhook
}

func isTokenValid(url string) bool {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("ERROR: '%s' is not a valid webhook URL.", url)
			os.Exit(0)
		}
	}()

	if url[0:33] == "https://discord.com/api/webhooks/" {
		url_r, err := requests.Get(url)
		manageError(err)

		url_code := url_r.StatusCode
		if url_code == 401 {
			return false
		} else {
			return true
		}
	} else {
		return false
	}
}

func isMsgMax(msg string) bool {
	msgLen := len(msg)
	msgLimit := 2000 // you never know if discord may change their
	// limit in the near future /shrug

	if msgLen < msgLimit {
		return false
	} else {
		msgToShort := msgLen - msgLimit
		fmt.Printf("Your message's length (%d) surpasses Discord's limit (%d)."+
			"Please make it %d characters shorter and try again.",
			msgLen, msgLimit, msgToShort)
		os.Exit(0)
	}
	return true
}

func manageError(err error) {
	// just to calm down with the syntax
	if err != nil {
		fmt.Println("An unexpected error ocurred. Please try again.")
		fmt.Println("For more information:")
		panic(err)
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
