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
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/jochasinga/requests"
	"github.com/spf13/cobra"
)

var get_cmd = &cobra.Command{
	Use:   "get [URL] [message-id]",
	Short: "Returns a previously-sent message (or an argument if a flag is used)",
	Args:  cobra.MinimumNArgs(2),
	// MaximumArgs: cobra.MaximumNArgs(3),
	// Long: maybe i won't use it, but i'll leave it here just in case

	Run: func(cmd *cobra.Command, args []string) {
		// i might move this to root.go
		flags := []bool{
			has_avatar_url, is_bot, discriminator, author_id, username_content,
			message_content, message_id, channel_id,
			mentions_everyone, mentions_roles, is_pinned, timestamp, has_tts,
			webhook_id, webhook_type,
			components, edited_timestamp, embeds, flags,
		}
		flags_var := []string{
			"avatar", "bot", "discriminator", "id", "username",
			"content", "id", "channel_id",
			"mention_everyone", "mention_roles", "pinned", "timestamp", "tts",
			"webhook_id", "wh_type",
			"components", "edited_timestamp", "embeds", "flags",
		}

		url := fmt.Sprintf("%s/messages/%s", args[0], args[1])
		if !is_token_valid(url) {
			fmt.Printf("ERROR: '%s' is not a valid webhook token", args[0])
		}

		resp_json, err := requests.Get(url)
		ManageError(err)
		var resp_map map[string]interface{}
		json.Unmarshal([]byte(resp_json.JSON()), &resp_map)

		// Checks for each flag
		var i int
		var isFlagSet bool
		for i = 0; i < len(flags); i++ {
			if flags[i] {
				if i >= 0 && i <= 4 {
					// checks if flag is on map:author
					resp_map_author := resp_map["author"]
					output := resp_map_author.(map[string]interface{})[flags_var[i]]
					output_type := reflect.TypeOf(output).String()
					if output_type == "bool" {
						output = fmt.Sprintf("%v", output)
					}

					fmt.Printf("%s: %s\n", flags_var[i], output)
					isFlagSet = true

				} else {
					output := resp_map[flags_var[i]]
					output_type := reflect.TypeOf(output).String()
					if output_type == "bool" {
						output = fmt.Sprintf("%v", output)
					}

					fmt.Printf("%s: %s\n", flags_var[i], output)
					isFlagSet = true
					// lol
				}
			} else {
				continue
			}
		}

		// default if no flag is set
		if !isFlagSet {
			output := resp_map["content"]
			fmt.Println(output)
		}
	},
}
