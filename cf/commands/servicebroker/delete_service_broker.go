package servicebroker

import (
	"github.com/cloudfoundry/cli/cf/api"
	"github.com/cloudfoundry/cli/cf/commandregistry"
	"github.com/cloudfoundry/cli/cf/configuration/coreconfig"
	"github.com/cloudfoundry/cli/cf/errors"
	"github.com/cloudfoundry/cli/cf/flags"
	. "github.com/cloudfoundry/cli/cf/i18n"
	"github.com/cloudfoundry/cli/cf/requirements"
	"github.com/cloudfoundry/cli/cf/terminal"
)

type DeleteServiceBroker struct {
	ui     terminal.UI
	config coreconfig.Reader
	repo   api.ServiceBrokerRepository
}

func init() {
	commandregistry.Register(&DeleteServiceBroker{})
}

func (cmd *DeleteServiceBroker) MetaData() commandregistry.CommandMetadata {
	fs := make(map[string]flags.FlagSet)
	fs["f"] = &flags.BoolFlag{ShortName: "f", Usage: T("Force deletion without confirmation")}

	return commandregistry.CommandMetadata{
		Name:        "delete-service-broker",
		Description: T("Delete a service broker"),
		Usage: []string{
			T("CF_NAME delete-service-broker SERVICE_BROKER [-f]"),
		},
		Flags: fs,
	}
}

func (cmd *DeleteServiceBroker) Requirements(requirementsFactory requirements.Factory, fc flags.FlagContext) []requirements.Requirement {
	if len(fc.Args()) != 1 {
		cmd.ui.Failed(T("Incorrect Usage. Requires an argument\n\n") + commandregistry.Commands.CommandUsage("delete-service-broker"))
	}

	reqs := []requirements.Requirement{
		requirementsFactory.NewLoginRequirement(),
	}

	return reqs
}

func (cmd *DeleteServiceBroker) SetDependency(deps commandregistry.Dependency, pluginCall bool) commandregistry.Command {
	cmd.ui = deps.UI
	cmd.config = deps.Config
	cmd.repo = deps.RepoLocator.GetServiceBrokerRepository()
	return cmd
}

func (cmd *DeleteServiceBroker) Execute(c flags.FlagContext) error {
	brokerName := c.Args()[0]
	if !c.Bool("f") && !cmd.ui.ConfirmDelete(T("service-broker"), brokerName) {
		return nil
	}

	cmd.ui.Say(T("Deleting service broker {{.Name}} as {{.Username}}...",
		map[string]interface{}{
			"Name":     terminal.EntityNameColor(brokerName),
			"Username": terminal.EntityNameColor(cmd.config.Username()),
		}))

	broker, err := cmd.repo.FindByName(brokerName)

	switch err.(type) {
	case nil:
	case *errors.ModelNotFoundError:
		cmd.ui.Ok()
		cmd.ui.Warn(T("Service Broker {{.Name}} does not exist.", map[string]interface{}{"Name": brokerName}))
		return nil
	default:
		return err
	}

	err = cmd.repo.Delete(broker.GUID)
	if err != nil {
		return err
	}

	cmd.ui.Ok()
	return nil
}
