package cmd

import (
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"

	"bitbucket.org/level27/lvl/types"
	"github.com/spf13/cobra"
)

var domainCmd = &cobra.Command{
	Use:   "domain",
	Short: "Commands for managing domains",
}

func init() {

	// domain command
	RootCmd.AddCommand(domainCmd)
	// Get (list of all domains)
	domainCmd.AddCommand(domainGetCmd)
	addCommonGetFlags(domainGetCmd)
	// Get details from a specific domain
	domainCmd.AddCommand(domainDescribeCmd)
	// Create (single domain)
	domainCmd.AddCommand(domainCreateCmd)
	
}




//GET LIST OF ALL DOMAINS [lvl domain get] 
var domainGetCmd = &cobra.Command{
	Use: "get",
	Short: "Get a list of all current domains",
	Run: func(ccmd *cobra.Command, args []string) {
		w := tabwriter.NewWriter(os.Stdout, 4, 8, 4, ' ', 0)
		fmt.Fprintln(w, "ID\tNAME\tSTATUS\t")

		domains := getDomains(args)
		for _, domain := range domains {
			fmt.Fprintln(w, strconv.Itoa(domain.ID)+"\t"+domain.Fullname+"\t"+domain.Status+"\t")
		}
	
		w.Flush()
	},
}


func getDomains(ids []string) []types.StructDomain {
	c := Level27Client
	if len(ids) == 0 {
		return c.Domains(optFilter, optNumber).Data
	} else {
		domains := make([]types.StructDomain, len(ids))
		for idx, id := range ids {
			domains[idx] = c.Domain("GET", id, nil).Data
		}
		return domains
	}
}


// CREATE DOMAIN [lvl domain get]
var domainCreateCmd = &cobra.Command{
	Use: "create",
	Short: "Create new a new domain",

	Run: func(ccmd *cobra.Command, args []string) {
		w := tabwriter.NewWriter(os.Stdout, 4, 8, 4, ' ', 0)
		fmt.Fprintln(w, "ID\tNAME\tSTATUS\t")

		domains := getDomains(args)
		for _, domain := range domains {
			fmt.Fprintln(w, strconv.Itoa(domain.ID)+"\t"+domain.Fullname+"\t"+domain.Status+"\t")
		}
	
		w.Flush()
	},
}


// DESCRIBE DOMAIN (get detailed info from specific domain) - [lvl domain describe <id>]
var domainDescribeCmd = &cobra.Command{
	Use:   "describe",
	Short: "Get detailed info about a domain",
	Run: func(cmd *cobra.Command, args []string) {
		Level27Client.DomainDescribe(args)
	},
}