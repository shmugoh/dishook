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
	"strconv"
	"strings"

	"github.com/jochasinga/requests"
	"github.com/spf13/cobra"
)

// flags
var (
	avatar_url string
	username   string
	message    string
	tts        bool
)

func init() {
	rootCmd.Flags().StringVarP(&avatar_url, "avatar-url", "a", "", "Sets the webhook's profile picture.")
	rootCmd.Flags().StringVarP(&message, "message", "m", "", "Sets the message you want to send.")
	rootCmd.Flags().StringVarP(&username, "username", "u", "", "Sets the username of the webhook.")
	rootCmd.Flags().BoolVarP(&tts, "tts", "t", false, "Sets if tts should be enabled or not.")
}

var rootCmd = &cobra.Command{
	Use:   " dishook [url] [message]\nor dishook [url]",
	Short: "CLI webhook runner for Discord.",
	Args:  cobra.ArbitraryArgs,
	// Args:  cobra.MinimumNArgs(2),
	// Long: maybe i won't use it, but i'll leave it here just in case

	Run: func(cmd *cobra.Command, args []string) {
		url := args[0]
		flags := []string{avatar_url, username, message}

		isTokenValid := isTokenValid(url)
		if isTokenValid == false {
			fmt.Printf("ERROR: '%s' is not a valid webhook token.", args[0])
		}

		// If flags are used
		for i := 0; i < len(flags); i++ {
			if len(flags[i]) != 0 {
				if len(message) == 0 {
					fmt.Printf("ERROR: Message flag required.")
					os.Exit(0)
				}

				tts := strconv.FormatBool(tts)
				values := map[string]string{
					"content":    message,
					"username":   username,
					"avatar_url": avatar_url,
					"tts":        tts,
				}

				jsonValue, _ := json.Marshal(values)
				fmt.Println(bytes.NewBuffer(jsonValue))
				resp, err := requests.Post(url, "application/json", bytes.NewBuffer(jsonValue))
				fmt.Print(resp)
				manageError(err)
				os.Exit(0)
			} else {
				continue
			}
		}

		// If flags ARE NOT used
		content := getContent(args, 1)
		sendMsg(url, content)
		// sendMsg(url, content)
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
