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

	"github.com/jochasinga/requests"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(getCmd)
}

var getCmd = &cobra.Command{
	Use:   "get [URL] [message-id]",
	Short: "Returns previously-sent webhook message as default (or one param if a flag is used).",
	Args:  cobra.MinimumNArgs(2),
	// MaximumArgs: cobra.MaximumNArgs(3),
	// Long: maybe i won't use it, but i'll leave it here just in case

	Run: func(cmd *cobra.Command, args []string) {
		url := args[0]
		message_id := args[1]
		url = url + "/messages/" + message_id

		isTokenValid := isTokenValid(url)
		if isTokenValid == false {
			fmt.Printf("ERROR: '%s' is not a valid webhook token.", args[0])
		}

		resp_json, err := requests.Get(url)
		manageError(err)

		// byteValue, _ := ioutil.ReadAll(resp_json)
		var resp_map map[string]interface{}
		json.Unmarshal([]byte(resp_json.JSON()), &resp_map)
		fmt.Println(resp_map["content"])
	},
}
