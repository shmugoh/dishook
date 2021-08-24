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

	"github.com/jochasinga/requests"
	"github.com/spf13/cobra"
)

func init() {
	editCmd.Flags().StringVarP(&message, "message", "m", "", "Sets the message you wanna edit.")

	rootCmd.AddCommand(editCmd)
}

var editCmd = &cobra.Command{
	Use:   "edit [URL] [message-id]",
	Short: "Edits a sent webhook message.",
	Args:  cobra.MinimumNArgs(2),
	// MaximumArgs: cobra.MaximumNArgs(3),
	// Long: maybe i won't use it, but i'll leave it here just in case

	Run: func(cmd *cobra.Command, args []string) {
		url := args[0]
		message_id := args[1]
		url = url + "/messages/" + message_id
		flags := []string{message}

		isTokenValid := isTokenValid(url)
		if isTokenValid == false {
			fmt.Printf("ERROR: '%s' is not a valid webhook token.", args[0])
		}

		for i := 0; i < len(flags); i++ { // checks if flags are used
			if len(flags[i]) != 0 {
				if len(message) == 0 {
					fmt.Printf("ERROR: Message flag required.")
					os.Exit(0)
				}

				values := map[string]string{
					"content": message,
				}

				jsonValue, _ := json.Marshal(values)
				fmt.Println(bytes.NewBuffer(jsonValue))
				resp, err := requests.Patch(url, "application/json", bytes.NewBuffer(jsonValue))
				fmt.Print(resp)
				manageError(err)
				os.Exit(0)
			} else {
				continue
			}
		}

		content := getContent(args, 2)
		values := map[string]string{"content": content}
		jsonValue, _ := json.Marshal(values)

		resp, err := requests.Patch(url, "application/json", bytes.NewBuffer(jsonValue))
		fmt.Print(resp)
		manageError(err)
	},
}
