// Copyright 2020 Layer5, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package system

import (
	"fmt"

	"github.com/layer5io/meshery/mesheryctl/internal/cli/root/system/context"

	"github.com/layer5io/meshery/mesheryctl/pkg/utils"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

const (
	url     = "http://localhost:9081"
	fileURL = "https://raw.githubusercontent.com/layer5io/meshery/master/docker-compose.yaml"
)

var (
	availableSubcommands = []*cobra.Command{}
)

// SystemCmd represents Meshery Lifecycle Management cli commands
var SystemCmd = &cobra.Command{
	Use:   "system",
	Short: "Meshery Lifecycle Management",
	Long:  `Manage the state and configuration of Meshery server, adapters, and client.`,
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if ok := utils.IsValidSubcommand(availableSubcommands, args[0]); !ok {
			return errors.New(utils.SystemError(fmt.Sprintf("invalid command: \"%s\"", args[0])))
		}
		return nil
	},
}

func init() {
	availableSubcommands = []*cobra.Command{
		resetCmd,
		logsCmd,
		startCmd,
		stopCmd,
		restartCmd,
		statusCmd,
		updateCmd,
		configCmd,
		context.ContextCmd,
		completionCmd,
		channelCmd,
	}
	SystemCmd.AddCommand(availableSubcommands...)
}
