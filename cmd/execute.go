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
	"strconv"

	"github.com/spf13/cobra"
)

// Automatically refers to execute()
var execute_cmd = &cobra.Command{
	Use:   "execute [URL] [message]\n  dishook execute [URL]",
	Short: "Sends message and/or arguments to Discord",
	Args:  cobra.MinimumNArgs(2),

	PreRunE: func(cmd *cobra.Command, args []string) error {
		url := args[0]
		if !is_token_valid(url) {
			return fmt.Errorf("'%s' not a valid webhook token", args[0])
		}
		return nil
	},

	RunE: func(cmd *cobra.Command, args []string) error {
		err := execute(args)
		if err != nil {
			ManageError(err)
		}
		return nil
	},
}

// Executes Discord Webhook.
//
// Available flags are "content", "username", "avatar_url" and "tts" (bool) and are parsed to a map[string].
//
// If no flags are parsed, execute will automatically resort to parse "content".
func execute(args []string) error {
	url := args[0]
	flags := []string{avatar_url, username, message}
	for i := 0; i < len(flags); i++ {
		if len(flags[i]) != 0 {
			if len(message) == 0 {
				return fmt.Errorf("message flag required")
			}
			if is_max(message) {
				return fmt.Errorf("message length surpasses 2000 character limit")
			}
			tts := strconv.FormatBool(tts)
			json_map := map[string]string{
				"content":    message,
				"username":   username,
				"avatar_url": avatar_url,
				"tts":        tts,
			}

			request_HTTP("POST", url, json_map)
			return nil
		}
	}

	// If flags are NOT used
	content := merge_strings(args, 1)
	if !is_max(content) {
		json_map := map[string]string{"content": content}
		request_HTTP("POST", url, json_map)
	} else {
		return fmt.Errorf("message length surpasses 2000 character limit")
	}
	return nil
}
