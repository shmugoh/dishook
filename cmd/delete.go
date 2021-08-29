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

	"github.com/jochasinga/requests"
	"github.com/spf13/cobra"
)

func init() {
	deleteCmd.SetUsageTemplate("\nUsage:\n  delete [URL] [message id]\n") // if i don't use this, usage would end with
	rootCmd.AddCommand(deleteCmd)                                         // [flags], and deletecmd doesn't use them
	// rootCmd.SetUsageTemplate("delete [URL] [message-id]")

}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Removes previously-sent webhook message",
	Args:  cobra.ExactArgs(2),

	Run: func(cmd *cobra.Command, args []string) {
		url := args[0]
		message_id := args[1]
		url = url + "/messages/" + message_id

		if !isTokenValid(url) {
			fmt.Printf("ERROR: '%s' is not a valid webhook token", args[0])
		}

		resp, err := requests.Delete(url)
		if err != nil || resp.StatusCode == 204 {
			fmt.Printf("Message with ID %s has been removed", message_id)
		} else {
			fmt.Printf("ERROR: '%s' message ID doesn't exist. Please check if message exists and try again", message_id)
		}
	},
}
