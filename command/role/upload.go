package rolecommand

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/go-chef/chef"
	"github.com/go-chef/gladius/app"
)

type UploadContext struct {
	log *logrus.Logger
	cfg *app.Configuration
}

func UploadCommand(env *app.Environment) cli.Command {
	c := &UploadContext{log: env.Log, cfg: env.Config}
	cmd := &cli.Command{
		Name:        "upload",
		Description: "Uploads the role(s) to the Chef server(s)",
		Usage:       "<role.json file(s)>",
		Action:      c.Run,
	}

	return *cmd
}

/*
 * upload <role.json>
 *
 */
func (r *UploadContext) Run(c *cli.Context) {
	if len(c.Args()) < 1 {
		cli.ShowCommandHelp(c, c.Command.Name)
		return
	}
	r.Do(c.Args())
}

func (r *UploadContext) Do(filenames []string) {
	log := r.log
	for _, chefServer := range r.cfg.ChefServers {
		for _, filename := range filenames {
			file, err := os.Open(filename)
			if err != nil {
				log.Errorln(fmt.Sprintf("Unable to open %s: %s", filename, err))
				continue
			}
			defer file.Close()

			v := &chef.Role{}
			err = json.NewDecoder(file).Decode(&v)
			if err != nil {
				log.Errorln(fmt.Sprintf("Invalid json in %s: %s", filename, err))
				continue
			}

			_, err = chefServer.Client.Roles.Get(v.Name)
			if err != nil {
				err = chefServer.Client.Roles.Create(v)
				if err != nil {
					log.Errorln(fmt.Sprintf("Error creating role from %s: %s", filename, err))
					continue
				}
				log.Infoln(fmt.Sprintf("Created the %s role on %s", v.Name, chefServer.ServerURL))
			} else {
				err = chefServer.Client.Roles.Put(v)
				if err != nil {
					log.Errorln(fmt.Sprintf("Error updating role from %s: %s", filename, err))
					continue
				}
				log.Infoln(fmt.Sprintf("Updated the %s role on %s", v.Name, chefServer.ServerURL))
			}
		}
	}
}
