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
	Short: "Returns info of a message sent by a webhook",
	Args:  cobra.MinimumNArgs(2),
	// MaximumArgs: cobra.MaximumNArgs(3),
	// Long: maybe i won't use it, but i'll leave it here just in case

	RunE: func(cmd *cobra.Command, args []string) error {
		url := fmt.Sprintf("%s/messages/%s", args[0], args[1])
		if !is_token_valid(url) {
			return fmt.Errorf("ERROR: '%s' is not a valid webhook token", args[0])
		}
		resp_json, err := requests.Get(url)
		ManageError(err)
		var resp_map map[string]interface{}
		json.Unmarshal([]byte(resp_json.JSON()), &resp_map)

		flags_map_bool := []bool{
			has_avatar_url, is_bot, discriminator, author_id, username_content,
			message_content, message_id, channel_id,
			mentions_everyone, mentions_roles, is_pinned, timestamp, has_tts,
			webhook_id, webhook_type,
			components, edited_timestamp, embeds, flags,
		}
		flags_map_var := []string{
			"avatar", "bot", "discriminator", "id", "username",
			"content", "id", "channel_id",
			"mention_everyone", "mention_roles", "pinned", "timestamp", "tts",
			"webhook_id", "wh_type",
			"components", "edited_timestamp", "embeds", "flags",
		}
		var is_flag_set bool

		for i := 0; i < len(flags_map_bool); i++ {
			if flags_map_bool[i] {
				output := resp_map[flags_map_var[i]]
				output_type := reflect.TypeOf(output).String()
				if output_type == "bool" {
					output = fmt.Sprintf("%v", output)
				}

				fmt.Printf("%s: %s\n", flags_map_var[i], output)
				is_flag_set = true // lol
			}
		}

		// default if no flag is set
		if !is_flag_set {
			fmt.Println(resp_map)
		}
		return nil
	},
}
