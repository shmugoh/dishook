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
	"log"
	"os"
	"strings"

	"github.com/jochasinga/requests"
	"github.com/spf13/cobra"
)

var (
	avatar_url string
	username   string
	message    string
	tts        bool

	author_id        bool
	username_content bool
	discriminator    bool
	components       bool
	message_content  bool
	channel_id       bool
	edited_timestamp bool
	embeds           bool
	flags            bool
	message_id       bool

	webhook_type bool
	webhook_id   bool
	timestamp    bool

	has_avatar_url    bool
	is_bot            bool
	mentions_everyone bool
	mentions_roles    bool
	is_pinned         bool
	has_tts           bool
)

func init() {
	// Execute Commands
	execute_cmd.Flags().StringVarP(&avatar_url, "avatar-url", "a", "", "Sets the webhook's profile picture")
	execute_cmd.Flags().StringVarP(&message, "message", "m", "", "Sets the message you wanna send")
	execute_cmd.Flags().StringVarP(&username, "username", "u", "", "Sets the username of the webhook")
	execute_cmd.Flags().BoolVarP(&tts, "tts", "t", false, "Sets if tts should be enabled or not")

	// Get Commands
	// json info | author:map
	get_cmd.Flags().BoolVarP(&has_avatar_url, "avatar-url", "a", false, "Avatar URL of the webhook")
	get_cmd.Flags().BoolVarP(&is_bot, "bot", "b", false, "Returns if obtained webhook is bot")
	get_cmd.Flags().BoolVarP(&discriminator, "discriminator", "d", false, "Webhook's discriminator")
	get_cmd.Flags().BoolVarP(&author_id, "author-id", "", false, "ID of webhook")
	get_cmd.Flags().BoolVarP(&username_content, "username", "u", false, "Name used for the webhook")

	// message info
	get_cmd.Flags().BoolVarP(&message_content, "message", "m", false, "Message sent (set to default if no flags are used)")
	get_cmd.Flags().BoolVarP(&message_id, "message-id", "s", false, "Message ID")
	get_cmd.Flags().BoolVarP(&channel_id, "channel-id", "c", false, "Channel ID")

	get_cmd.Flags().BoolVarP(&mentions_everyone, "mentions-everyone", "e", false, "Returns if everyone is mentioned")
	get_cmd.Flags().BoolVarP(&mentions_roles, "mention-roles", "r", false, "Returns if roles are mentioned")
	get_cmd.Flags().BoolVarP(&is_pinned, "pinned", "p", false, "Returns if message is pinned")
	get_cmd.Flags().BoolVarP(&timestamp, "timestamp", "", false, "Returns the time message was sent")
	get_cmd.Flags().BoolVarP(&has_tts, "tts", "t", false, "Returns if TTS was used")

	// webhook info
	get_cmd.Flags().BoolVarP(&webhook_id, "webhook-id", "", false, "Webhook's ID")
	get_cmd.Flags().BoolVarP(&webhook_type, "webhook-type", "", false, "Webhook type")

	// misc
	get_cmd.Flags().BoolVarP(&components, "components", "", false, "Components included with the message")
	get_cmd.Flags().BoolVarP(&edited_timestamp, "edited-timestamp", "", false, "Time when message was edited")
	get_cmd.Flags().BoolVarP(&embeds, "embeds", "", false, "Array of message embeds/components")
	get_cmd.Flags().BoolVarP(&flags, "flags", "", false, "Name of the webhook")

	// Edit Command (lol)
	edit_cmd.Flags().StringVarP(&message, "message", "m", "", "Sets the message you wanna edit")

	root_cmd.AddCommand(get_cmd, execute_cmd, edit_cmd, delete_cmd)
}

// Allocates all strings from array starting from
// given argument position into one variable with
// blank spaces
//
// Used when input has no flags set.
func merge_strings(args []string, arg_pos int) string {
	var str string
	for i := arg_pos; i < len(args); i++ {
		str = fmt.Sprintf("%s %s", str, args[i])
	}
	return strings.TrimSpace(str)
}

// Semds an HTTP request method to Discord's webhook
// with provided URL and JSON map.
// Automatically marshalls the JSON map.
//
// Supported HTTP Methods: POST, PATCH
func request_HTTP(http_method string, URL string, json_map map[string]string) {
	json_value, _ := json.Marshal(json_map)

	switch http_method {
	case "POST":
		resp, err := requests.Post(URL, "application/json", bytes.NewBuffer(json_value))
		_, _ = resp, err
	case "PATCH":
		resp, err := requests.Patch(URL, "application/json", bytes.NewBuffer(json_value))
		_, _ = resp, err
	}
}

// Checks if provided URL matches the Discord URL webhook API calling one.
func is_token_valid(url string) bool {
	if url[0:33] == "https://discord.com/api/webhooks/" {
		url_r, err := requests.Get(url)
		url_code := url_r.StatusCode
		if url_code != 401 || err != nil {
			return true
		}
	}

	// Passed thru all checks without any true statement
	return false

	// defer func() {
	// 	if err := recover(); err != nil {
	// 		// what
	// 		// i might tickle around and see what this does
	// 		fmt.Printf("ERROR: '%s' is not a valid webhook URL.", url)
	// 		os.Exit(0)
	// 	}
	// }()
}

// Checks if given value doesn't pass set (2000) characters.
func is_max(msg string) bool {
	msg_limit := 2000
	return len(msg) >= msg_limit
}

// Panics when an error ocurrs. Only use if no conditionals are being used or if it's a HTTPS request.
func ManageError(err error) {
	if err != nil {
		log.Fatal(err)
		fmt.Println("An unexpected error ocurred. Please try again.")
	}
}

// Cobra stuff
var root_cmd = &cobra.Command{
	Use:  "dishook [url] [message]\n  dishook [url] [flags]",
	Args: cobra.MinimumNArgs(2),

	PreRunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			cmd.Help()
		}
		return nil
	},

	Run: func(cmd *cobra.Command, args []string) {
		execute(args)
	},
}

func Execute() {
	if err := root_cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
