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
	"os/exec"

	"github.com/pkg/errors"

	"github.com/layer5io/meshery/mesheryctl/internal/cli/root/config"
	"github.com/layer5io/meshery/mesheryctl/pkg/utils"

	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// dashboardCmd represents the - command
var dashboardCmd = &cobra.Command{
	Use:   "dashboard",
	Short: "Open Meshery UI",
	Long:  `Well we might change it later`,
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := dashboard(); err != nil {
			return errors.Wrap(err, "failed to open the Meshery UI")

		}
		return nil

	},
}

func dashboard() error {

	// Get viper instance used for context
	mctlCfg, err := config.GetMesheryCtl(viper.GetViper())
	if err != nil {
		return errors.Wrap(err, "error processing config")
	}
	// get the platform, channel and the version of the current context
	// if a temp context is set using the -c flag, use it as the current context
	currCtx, err := mctlCfg.SetCurrentContext(tempContext)
	if err != nil {
		return err
	}
	//Adding a timeout for a 2 second delay
	timeout := time.Duration(1 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	//Testing if the UI is accesible at the endpoint
	resp, err := client.Get(currCtx.Endpoint)
	if err != nil {
		return err
	}

	currPlatform := currCtx.Platform

	err = utils.NavigateToBrowser(currCtx.Endpoint)

	if err != nil {

		switch currPlatform {
		case "docker":
			//Nothing can be done ?
			break
		case "kubernetes":
			// Probably works
			exec.Command("kubectl proxy")

		}

		err = utils.NavigateToBrowser(currCtx.Endpoint)
		if err != nil {
			return err
		}

	}

	log.Info("Meshery UI is accesible at " + currCtx.Endpoint + "status:" + string(rune(resp.StatusCode)))
	log.Info("Opening Meshery in your browser. If Meshery does not open, please point your browser to " + currCtx.Endpoint + " to access Meshery.")

	return nil
}
