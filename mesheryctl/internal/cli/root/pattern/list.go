package pattern

import (
	"encoding/json"
<<<<<<< HEAD
<<<<<<< HEAD
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"
=======
	"net/http"
>>>>>>> 306424b1... list command v1
=======
	"io/ioutil"
	"net/http"
	"time"
>>>>>>> 2c41dd5c... wip list command

	"github.com/layer5io/meshery/mesheryctl/internal/cli/root/config"
	"github.com/layer5io/meshery/mesheryctl/pkg/utils"
	"github.com/layer5io/meshery/models"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
<<<<<<< HEAD
<<<<<<< HEAD
	token   string
	verbose bool
)

var listCmd = &cobra.Command{
	Use:  "list",
	Long: "Display list of all available pattern files",
	Args: cobra.MinimumNArgs(0),
=======
	token string
=======
	token   string
	allflag bool
>>>>>>> 2c41dd5c... wip list command
)
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list pattern files",
<<<<<<< HEAD
	Long:  "List available pattern files",
	Args:  cobra.MinimumNArgs(1),
>>>>>>> 306424b1... list command v1
=======
	Long:  "Display list of all available pattern files",
	Args:  cobra.MinimumNArgs(0),
>>>>>>> 2c41dd5c... wip list command
	RunE: func(cmd *cobra.Command, args []string) error {
		mctlCfg, err := config.GetMesheryCtl(viper.GetViper())
		if err != nil {
			return errors.Wrap(err, "error processing config")
		}
<<<<<<< HEAD
<<<<<<< HEAD
		var response models.PatternsAPIResponse

		client := &http.Client{}
		req, err := http.NewRequest("GET", mctlCfg.GetBaseMesheryURL()+"/api/experimental/patternfile", nil)
		if err != nil {
			return err
		}
		err = utils.AddAuthDetails(req, token)
		if err != nil {
			return err
		}
		res, err := client.Do(req)
		if err != nil {
			return err
		}
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}
		err = json.Unmarshal(body, &response)
=======
		var Patterns models.PatternsApiResponse
=======
		var response models.PatternsApiResponse
>>>>>>> 2c41dd5c... wip list command

		client := &http.Client{}
		req, err := http.NewRequest("GET", mctlCfg.GetBaseMesheryURL()+"/api/experimental/patternfile", nil)
		if err != nil {
			return err
		}
		err = utils.AddAuthDetails(req, token)
		if err != nil {
			return err
		}
<<<<<<< HEAD
		client.Do(req)
		err = json.NewDecoder(req.Body).Decode(&Patterns)
>>>>>>> 306424b1... list command v1
		if err != nil {
			return err
		}

<<<<<<< HEAD
		tokenObj, err := utils.ReadToken(token)
		if err != nil {
			return err
		}
		provider := tokenObj["meshery-provider"]
		var data [][]string

		if verbose == true {
			if provider == "None" {
				for _, v := range response.Patterns {
					PatternID := v.ID.String()
					PatterName := v.Name
					CreatedAt := fmt.Sprintf("%d-%d-%d %d:%d:%d", int(v.CreatedAt.Month()), v.CreatedAt.Day(), v.CreatedAt.Year(), v.CreatedAt.Hour(), v.CreatedAt.Minute(), v.CreatedAt.Second())
					UpdatedAt := fmt.Sprintf("%d-%d-%d %d:%d:%d", int(v.UpdatedAt.Month()), v.UpdatedAt.Day(), v.UpdatedAt.Year(), v.UpdatedAt.Hour(), v.UpdatedAt.Minute(), v.UpdatedAt.Second())
					data = append(data, []string{PatternID, PatterName, CreatedAt, UpdatedAt})
				}
				utils.PrintToTable([]string{"PATTERN ID", "NAME", "CREATED", "UPDATED"}, data)
				return nil
			}
			for _, v := range response.Patterns {
				PatternID := utils.TruncateID(v.ID.String())
				var UserID string
				if v.UserID != nil {
					UserID = *v.UserID
				} else {
					UserID = "null"
				}
				PatterName := v.Name
				CreatedAt := fmt.Sprintf("%d-%d-%d", int(v.CreatedAt.Month()), v.CreatedAt.Day(), v.CreatedAt.Year())
				UpdatedAt := fmt.Sprintf("%d-%d-%d", int(v.UpdatedAt.Month()), v.UpdatedAt.Day(), v.UpdatedAt.Year())
				data = append(data, []string{PatternID, UserID, PatterName, CreatedAt, UpdatedAt})
			}
			utils.PrintToTable([]string{"PATTERN ID", "USER ID", "NAME", "CREATED", "UPDATED"}, data)

			return nil
		}

		// Check if messhery provider is set
		if provider == "None" {
			for _, v := range response.Patterns {
				PatterName := fmt.Sprintf("%s", strings.Trim(v.Name, filepath.Ext(v.Name)))
				PatternID := utils.TruncateID(v.ID.String())
				CreatedAt := fmt.Sprintf("%d-%d-%d", int(v.CreatedAt.Month()), v.CreatedAt.Day(), v.CreatedAt.Year())
				UpdatedAt := fmt.Sprintf("%d-%d-%d", int(v.UpdatedAt.Month()), v.UpdatedAt.Day(), v.UpdatedAt.Year())
				data = append(data, []string{PatternID, PatterName, CreatedAt, UpdatedAt})
			}
			utils.PrintToTable([]string{"PATTERN ID", "NAME", "CREATED", "UPDATED"}, data)
		}
		return nil
	},
}

func init() {
	listCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Display full length user and pattern file identifiers")
	listCmd.Flags().StringVarP(&token, "token", "t", "", "path to token")
	_ = listCmd.MarkFlagRequired("token")
}
=======
	},
}
>>>>>>> 306424b1... list command v1
=======
		res, err := client.Do(req)
		if err != nil {
			return err
		}
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}
		json.Unmarshal(body, &response)
		if allflag == true {
			for _, v := range response.Patterns {
				headers := []string{"PATTERN ID", "NAME", "USER ID", "CREATED", "UPDATED"}
				data := [][]string{
					{v.ID.String(), v.Name, v.ID.String(), v.CreatedAt.Format(time.RFC3339Nano), v.UpdatedAt.Format(time.RFC3339Nano)},
				}
				utils.PrintToTable(headers, data)
			}
			return nil
		}
		for _, v := range response.Patterns {
			headers := []string{"NAME", "USER ID", "CREATED", "UPDATED"}
			data := [][]string{
				{v.Name, v.ID.String(), v.CreatedAt.Format(time.RFC3339Nano), v.UpdatedAt.Format(time.RFC3339Nano)},
			}
			utils.PrintToTable(headers, data)
		}
		return nil
	},
}

func init() {
	listCmd.Flags().BoolVarP(&allflag, "all", "a", false, "Display full length user and pattern file identifiers")
	listCmd.MarkFlagRequired("token")
}
>>>>>>> 2c41dd5c... wip list command
