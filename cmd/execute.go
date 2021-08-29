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
	"fmt"
	"os"
	"strconv"

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
	rootCmd.Flags().StringVarP(&avatar_url, "avatar-url", "a", "", "Sets the webhook's profile picture")
	rootCmd.Flags().StringVarP(&message, "message", "m", "", "Sets the message you wanna send")
	rootCmd.Flags().StringVarP(&username, "username", "u", "", "Sets the username of the webhook")
	rootCmd.Flags().BoolVarP(&tts, "tts", "t", false, "Sets if tts should be enabled or not")
	rootCmd.AddCommand(executeCmd)
}

var executeCmd = &cobra.Command{
	Use:   "execute [URL] [message]\n  dishook execute [URL]",
	Short: "Sends the message (with its corresponding flags if called)",
	Args:  cobra.MinimumNArgs(2),

	Run: func(cmd *cobra.Command, args []string) {
		execute(args)
	},
}

func execute(args []string) {
	url := args[0]
	flags := []string{avatar_url, username, message}
	if !isTokenValid(url) {
		fmt.Printf("ERROR: '%s' is not a valid webhook token.", args[0])
	}

	// Flags zone
	for i := 0; i < len(flags); i++ {
		if len(flags[i]) != 0 {
			if len(message) == 0 {
				fmt.Println("ERROR: Message flag required.")
				os.Exit(0)
			}
			if isMsgMax(message) {
				fmt.Println("ERROR: Message's length surpasses 2000 characters." +
					"Please make it shorter and try again.")
				os.Exit(0)
			}
			// TODO: create a handling error script
			// the one golang provides isn't good ux-wise

			tts := strconv.FormatBool(tts)
			json_map := map[string]string{
				"content":    message,
				"username":   username,
				"avatar_url": avatar_url,
				"tts":        tts,
			}
			requestHTTP("POST", url, json_map)
		} else {
			continue
		}
	}

	// No flags zone
	content := mergeStrings(args, 1)
	if isMsgMax(content) {
		fmt.Println("Your message's length surpasses 2000 characters." +
			"Please make it shorter and try again.")
	} else {
		json_map := map[string]string{"content": content}
		requestHTTP("POST", url, json_map)
	}

	// sendMsg(url, content)
}
