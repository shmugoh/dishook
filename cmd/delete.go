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
	delete_cmd.SetUsageTemplate("\nUsage:\n  delete [URL] [message id]\n")
}

// Deletes Discord Webhook.
//
// Required arguments are the Webhook's token and message ID.
var delete_cmd = &cobra.Command{
	Use:   "delete",
	Short: "Removes sent webhook message",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		url := fmt.Sprintf("%s/messages/%s", args[0], args[1])
		if !is_token_valid(url) {
			ManageError(fmt.Errorf("'%s' not a valid webhook token", args[0]))
		}

		resp, err := requests.Delete(url)
		if resp.StatusCode == 204 || err != nil {
			fmt.Printf("Message ID %s removed", args[1])
		} else {
			ManageError(fmt.Errorf("'%s' message ID not found", args[1]))
		}
		return nil
	},
}
