package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/level27/l27-go"
	"github.com/spf13/cobra"
)

var organisationCmd = &cobra.Command{
	Use:   "organisation",
	Short: "Commands for managing organisations",
}

var organisationGetCmd = &cobra.Command{
	Use: "get",

	Args: cobra.ArbitraryArgs,
	Run: func(ccmd *cobra.Command, args []string) {
		ids, err := convertStringsToIds(args)
		if err != nil {
			log.Fatalln("Invalid organisation ID")
		}
		outputFormatTable(getOrganisations(ids), []string{"ID", "NAME"}, []string{"ID", "Name"})
	},
}

func init() {
	RootCmd.AddCommand(organisationCmd)

	organisationCmd.AddCommand(organisationGetCmd)
	addCommonGetFlags(organisationGetCmd)
}

func resolveOrganisation(arg string) int {
	id, err := strconv.Atoi(arg)
	if err == nil {
		return id
	}

	return resolveShared(
		Level27Client.LookupOrganisation(arg),
		arg,
		"organisation",
		func(app l27.Organisation) string { return fmt.Sprintf("%s (%d)", app.Name, app.ID) }).ID
}

func getOrganisations(ids []int) []l27.Organisation {
	c := Level27Client
	if len(ids) == 0 {
		return c.Organisations(optGetParameters)
	} else {
		organisations := make([]l27.Organisation, len(ids))
		for idx, id := range ids {
			organisations[idx] = c.Organisation(id)
		}
		return organisations
	}
}
