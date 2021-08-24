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
	rootCmd.AddCommand(executeCmd)
}

var executeCmd = &cobra.Command{
	Use:   "execute [URL] [message]",
	Short: "Executor lol.",
	Args:  cobra.MinimumNArgs(2),

	Run: func(cmd *cobra.Command, args []string) {
		execute(args)
	},
}

func execute(args []string) {
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
}
