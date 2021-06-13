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
	"os"
	"reflect"

	"github.com/jochasinga/requests"
	"github.com/spf13/cobra"
)

// Flags
var (
	// author:map
	hasAvatarUrl     bool
	isBot            bool
	discriminator    bool
	author_id        bool
	username_content bool

	channel_id       bool
	components       bool
	message_content  bool
	edited_timestamp bool
	embeds           bool
	flags            bool
	message_id       bool
	mentionsEveryone bool
	mentionsRoles    bool
	isPinned         bool
	timestamp        bool
	hasTTS           bool
	wh_type          bool
	webhook_id       bool
)

func init() {
	// author:map
	getCmd.Flags().BoolVarP(&hasAvatarUrl, "avatar-url", "a", false, "Returns obtained webhook's avatar URL.")
	getCmd.Flags().BoolVarP(&isBot, "bot", "b", false, "Returns if obtained webhook is bot. (lol)") // you gotta make it detailed for a cli app!
	getCmd.Flags().BoolVarP(&discriminator, "discriminator", "d", false, "Returns webhook's discriminator.")
	getCmd.Flags().BoolVarP(&author_id, "author-id", "", false, "Returns the id of the webhook.")
	getCmd.Flags().BoolVarP(&username_content, "username", "u", false, "Returns the name used for the webhook.")

	// message info
	getCmd.Flags().BoolVarP(&message_content, "message", "m", false, "Returns the message. (set to default if no flags are used.)")
	getCmd.Flags().BoolVarP(&message_id, "message-id", "s", false, "Returns the message id.")
	getCmd.Flags().BoolVarP(&channel_id, "channel-id", "c", false, "Returns the channel id.")

	getCmd.Flags().BoolVarP(&mentionsEveryone, "mentions-everyone", "e", false, "Returns if everyone is mentioned.")
	getCmd.Flags().BoolVarP(&mentionsRoles, "mention-roles", "r", false, "Returns if roles are mentioned")
	getCmd.Flags().BoolVarP(&isPinned, "pinned", "p", false, "Returns if message is pinned.")
	getCmd.Flags().BoolVarP(&timestamp, "timestamp", "", false, "Returns the time message was sent.")
	getCmd.Flags().BoolVarP(&hasTTS, "tts", "t", false, "Returns if TTS was used.")

	// webhook details
	getCmd.Flags().BoolVarP(&webhook_id, "webhook-id", "", false, "Returns the webhook's id.")
	getCmd.Flags().BoolVarP(&wh_type, "webhook-type", "", false, "Returns the webhook type.")

	// misc
	getCmd.Flags().BoolVarP(&components, "components", "", false, "Returns the components included with the message.")
	getCmd.Flags().BoolVarP(&edited_timestamp, "edited-timestamp", "", false, "Returns the time the message was edited.")
	getCmd.Flags().BoolVarP(&embeds, "embeds", "", false, "Returns the array of message components.")
	getCmd.Flags().BoolVarP(&flags, "flags", "", false, "Returns the name used for the webhook")

	// what the hell

	rootCmd.AddCommand(getCmd)
}

var getCmd = &cobra.Command{
	Use:   "get [URL] [message-id]",
	Short: "Returns previously-sent webhook message as default (or one param if a flag is used).",
	Args:  cobra.MinimumNArgs(2),
	// MaximumArgs: cobra.MaximumNArgs(3),
	// Long: maybe i won't use it, but i'll leave it here just in case

	Run: func(cmd *cobra.Command, args []string) {
		flags := []bool{
			hasAvatarUrl, isBot, discriminator, author_id, username_content,
			message_content, message_id, channel_id,
			mentionsEveryone, mentionsRoles, isPinned, timestamp, hasTTS,
			webhook_id, wh_type,
			components, edited_timestamp, embeds, flags,
		}
		flags_var := []string{
			"avatar", "bot", "discriminator", "id", "username",
			"content", "id", "channel_id",
			"mention_everyone", "mention_roles", "pinned", "timestamp", "tts",
			"webhook_id", "wh_type",
			"components", "edited_timestamp", "embeds", "flags",
		}

		url := args[0]
		message_id := args[1]
		url = url + "/messages/" + message_id

		isTokenValid := isTokenValid(url)
		if isTokenValid == false {
			fmt.Printf("ERROR: '%s' is not a valid webhook token.", args[0])
		}

		resp_json, err := requests.Get(url)
		manageError(err)
		var resp_map map[string]interface{}
		json.Unmarshal([]byte(resp_json.JSON()), &resp_map)

		// Flags
		var i int
		var isFlagSet bool
		for i = 0; i < len(flags); i++ {
			if flags[i] == true {

				if i >= 0 && i <= 4 {
					// checks if flag is on map:author
					resp_map_author := resp_map["author"]
					output := resp_map_author.(map[string]interface{})[flags_var[i]]
					output_type := reflect.TypeOf(output).String()
					if output_type == "bool" {
						output = fmt.Sprintf("%v", output)
					}
					fmt.Printf("%s: %s\n", flags_var[i], output)

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
		if isFlagSet == false {
			output := resp_map["content"]
			fmt.Println(output)
		} else { // unless...
			os.Exit(0)
		}
	},
}
