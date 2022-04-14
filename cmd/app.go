package cmd

import (
	"fmt"
	"log"
	"strconv"

	"bitbucket.org/level27/lvl/types"
	"bitbucket.org/level27/lvl/utils"
	"github.com/spf13/cobra"
)

func init() {
	// ---- MAIN COMMAND APP
	RootCmd.AddCommand(appCmd)

	//------------------------------------------------- APP (GET / CREATE / DELETE / UPDATE / DESCRIBE)-------------------------------------------------
	// #region APP MAIN COMMANDS (GET / CREATE / UPDATE / DELETE / DESCRIBE)

	// ---- GET
	appCmd.AddCommand(appGetCmd)
	addCommonGetFlags(appGetCmd)

	// ---- CREATE
	appCmd.AddCommand(appCreateCmd)
	// flags used for creating app
	flags := appCreateCmd.Flags()
	flags.StringVarP(&appCreateName, "name", "n", "", "Name of the app.")
	flags.StringVarP(&appCreateOrg, "organisation", "", "", "The name of the organisation/owner of the app.")
	flags.IntSliceVar(&appCreateTeams, "autoTeams", appCreateTeams, "A csv list of team ID's.")
	flags.StringVar(&appCreateExtInfo, "externalInfo", "", "ExternalInfo (required when billableItemInfo entities for an organisation exist in DB.)")
	appCreateCmd.MarkFlagRequired("name")
	appCreateCmd.MarkFlagRequired("organisation")

	// ---- DELETE APP
	appCmd.AddCommand(appDeleteCmd)
	//flag to skip confirmation when deleting an app
	appDeleteCmd.Flags().BoolVarP(&isAppDeleteConfirmed, "yes", "y", false, "Set this flag to skip confirmation when deleting an app")

	// ---- UPDATE APP
	appCmd.AddCommand(appUpdateCmd)
	// flags needed for update command
	flags = appUpdateCmd.Flags()
	flags.StringVarP(&appUpdateName, "name", "n", "", "Name of the app.")
	flags.StringVarP(&appUpdateOrg, "organisation", "", "", "The name of the organisation/owner of the app.")
	flags.StringSliceVar(&appUpdateTeams, "autoTeams", appUpdateTeams, "A csv list of team ID's.")

	// ---- DESCRIBE APP
	appCmd.AddCommand(AppDescribeCmd)
	// #endregion

	//-------------------------------------------------  APP ACTIONS (ACTIVATE / DEACTIVATE) -------------------------------------------------
	// #region  APP ACTIONS (ACTIVATE / DEACTIVATE)

	// ACTION COMMAND
	appCmd.AddCommand(AppActionCmd)

	// ACTIVATE APP
	AppActionCmd.AddCommand(AppActionActivateCmd)

	// DEACTIVATE APP
	AppActionCmd.AddCommand(AppActionDeactivateCmd)
	// #endregion

	//------------------------------------------------- APP COMPONENTS (CREATE / GET / UPDATE / DELETE / DESCRIBE)-------------------------------------------------
	// ----	COMPONENT COMMAND
	appCmd.AddCommand(appComponentCmd)

	// ---- GET COMPONENTS
	appComponentCmd.AddCommand(appComponentGetCmd)
	addCommonGetFlags(appComponentGetCmd)

	// ---- CREATE COMPONENT
	appComponentCmd.AddCommand(appComponentCreateCmd)

	//------------------------------------------------- APP COMPONENTS HELPERS (CATEGORY )-------------------------------------------------
	// ---- GET COMPONENT CATEGORIES 
	appComponentCmd.AddCommand(appComponentCategoryGetCmd)

	// ---- GET COMPONENTTYPES
	appComponentCmd.AddCommand(appComponentTypeCmd)

	//-------------------------------------------------  APP SSL CERTIFICATES (GET / CREATE) -------------------------------------------------
	appCmd.AddCommand(appCertificateCmd)

	// ---- GET SSL CERTIFICATES 
	appCertificateCmd.AddCommand(appCertificateGetCmd)

	// ---- ADD SSL CERTIFICATE
	appCertificateCmd.AddCommand(appCertificateAddCmd)
	// flags used for adding a certificate to an app
	flags = appCertificateAddCmd.Flags()
	flags.StringVarP(&appAddSslName, "name", "n", "", "The name of the certificate.")
	flags.StringVarP(&appAddSslType, "type", "t", "", "The type of the certificate.")
	flags.StringArrayVarP(&appAddSslAutoUrl, "certificateUrl", "c", []string{}, "AutoSslCertificateUrls: url or csv list of urls (required for type letsencrypt).")
	flags.StringVarP(&appAddSslKey, "key", "k", "", "Ssl key (required for ssl type own).")
	flags.StringVar(&appAddSslCrt, "crt", "", "Ssl crt (Required for ssl type own).")
	flags.StringVar(&appAddSslCabundle, "cabundle", "", "Ssl cabundle (Required for ssl type own)." )
	flags.BoolVar(&appAddSslAutoUrlLink, "autoUrl", false, "If 'autoUrlLink' is set to true then a certificate's urls, which don't have another cettificate, will be linked to the certificate after successful creation (default: false).")
	flags.BoolVar(&appAddSslForce, "force", true, "Force ssl (default: true).")
}

// MAIN COMMAND APPS
var appCmd = &cobra.Command{
	Use:     "app",
	Short:   "Commands to manage apps",
	Example: "lvl app [subcommmand]",
}

//------------------------------------------------- APP (GET / CREATE  / UPDATE / DELETE / DESCRIBE)-------------------------------------------------
// #region APP MAIN SUBCOMMANDS (GET / CREATE / UPDATE / DELETE / DESCRIBE)

// ---- GET apps
var appGetCmd = &cobra.Command{
	Use:     "get",
	Short:   "Shows a list of all available apps.",
	Example: "lvl app get",
	Args:    cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		ids, err := convertStringsToIds(args)
		if err != nil {
			log.Fatalln("Invalid app ID")
		}

		outputFormatTable(
			getApps(ids),
			[]string{"ID", "NAME", "STATUS"},
			[]string{"ID", "Name", "Status"})
	},
}

func getApps(ids []int) []types.App {
	c := Level27Client
	if len(ids) == 0 {
		return c.Apps(optGetParameters)
	} else {
		apps := make([]types.App, len(ids))
		for idx, id := range ids {
			apps[idx] = c.App(id)
		}
		return apps
	}
}

// ---- CREATE NEW APP
var appCreateName, appCreateOrg, appCreateExtInfo string
var appCreateTeams []int
var appCreateCmd = &cobra.Command{
	Use:     "create",
	Short:   "Create a new app.",
	Example: "lvl app create -n myNewApp --organisation level27",
	Run: func(cmd *cobra.Command, args []string) {
		// check if name is valid.
		if appCreateName == "" {
			log.Fatalln("app name cannot be empty.")
		}

		// fill in all the props needed for the post request
		organisation := resolveOrganisation(appCreateOrg)
		request := types.AppPostRequest{
			Name:         appCreateName,
			Organisation: organisation,
			AutoTeams:    appCreateTeams,
			ExternalInfo: appCreateExtInfo,
		}

		// when succesfully creating app. app will be returned
		app := Level27Client.AppCreate(request)
		log.Printf("Succesfully created app! [name: '%v' - ID: '%v']", app.Name, app.ID)

	},
}

// ---- DELETE AN APP
var isAppDeleteConfirmed bool
var appDeleteCmd = &cobra.Command{
	Use:     "delete",
	Short:   "Delete an app",
	Example: "lvl app delete 2593",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// check for valid appID type
		appId := checkSingleIntID(args[0], "app")

		Level27Client.AppDelete(appId, isAppDeleteConfirmed)
	},
}

// ---- UPDATE AN APP
var appUpdateName, appUpdateOrg string
var appUpdateTeams []string
var appUpdateCmd = &cobra.Command{
	Use:     "update [appID]",
	Short:   "Update an app.",
	Example: "lvl app update 2067 --name myUpdatedName",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		//check if appId is valid
		appId := checkSingleIntID(args[0], "app")

		//get the current data from the app. if not changed its needed for put request
		currentData := Level27Client.App(appId)

		var currentTeamIds []string
		for _, team := range currentData.Teams {
			currentTeamIds = append(currentTeamIds, strconv.Itoa(team.ID))
		}
		// fill in request with the current data.
		request := types.AppPutRequest{
			Name:         currentData.Name,
			Organisation: currentData.Organisation.ID,
			AutoTeams:    currentTeamIds,
		}

		//when flags have been set. we need the currentdata to be updated.
		if cmd.Flag("name").Changed {
			request.Name = appUpdateName
		}

		if cmd.Flag("organisation").Changed {
			organisationID := resolveOrganisation(appUpdateOrg)
			request.Organisation = organisationID
		}

		if cmd.Flag("autoTeams").Changed {
			request.AutoTeams = appUpdateTeams
		}

		Level27Client.AppUpdate(appId, request)

	},
}

// ---- DESCRIBE APP
var AppDescribeCmd = &cobra.Command{
	Use:     "describe",
	Short:   "Get detailed info about an app.",
	Example: "lvl app describe 2077",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		//check for valid appId
		appId := checkSingleIntID(args[0], "app")
		// get all data from app by appId
		app := Level27Client.App(appId)
		outputFormatTemplate(app, "templates/app.tmpl")
	},
}

// #endregion

//------------------------------------------------- APP ACTIONS (ACTIVATE / DEACTIVATE)-------------------------------------------------
// #region APP ACTIONS (ACTIVATE / DEACTIVATE)

// ---- ACTION COMMAND
var AppActionCmd = &cobra.Command{
	Use:     "action",
	Short:   "commands to run specific actions on an app",
	Example: "lvl app action [subcommand]",
}

// ---- ACTIVATE APP
var AppActionActivateCmd = &cobra.Command{
	Use:     "activate",
	Short:   "Activate an app",
	Example: "lvl app action activate 2077",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// check for valid appId
		appId := checkSingleIntID(args[0], "app")

		Level27Client.AppAction(appId, "activate")
	},
}

// ---- DEACTIVATE APP
var AppActionDeactivateCmd = &cobra.Command{
	Use:     "deactivate",
	Short:   "Deactivate an app",
	Example: "lvl app action deactivate 2077",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// check for valid appId
		appId := checkSingleIntID(args[0], "app")

		Level27Client.AppAction(appId, "deactivate")
	},
}

// #endregion

//------------------------------------------------- APP COMPONENTS (CREATE / GET / UPDATE / DELETE / DESCRIBE)-------------------------------------------------
// ---- COMPONENT COMMAND
var appComponentCmd = &cobra.Command{
	Use: "component",
	Short: "Commands for managing appcomponents.",
	Example: "lvl app component get",
}

// ---- GET COMPONENTS
var appComponentGetCmd = &cobra.Command{
	Use: "get",
	Short: "Show list of all available components",
	Example: "lvl app component get",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		
		//check for valid appId
		appId := checkSingleIntID(args[0], "apps")
		ids, err := convertStringsToIds(args[1:])
		if err != nil {
			log.Fatalln("Invalid component ID")
		}
		log.Print(ids)
		outputFormatTable(
			getComponents(appId, ids),
			[]string{"ID", "NAME", "STATUS"},
			[]string{"ID", "Name", "Status"})
	},
}

func getComponents(appId int ,ids []int) []types.AppComponent2 {
	c := Level27Client
	if len(ids) == 0 {
		return c.AppComponentsGet(appId,optGetParameters)
	} else {
		components := make([]types.AppComponent2, len(ids))
		for idx, id := range ids {
			components[idx] = c.AppComponentGetSingle(appId, id)
		}
		return components
	}
}


// ---- CREATE COMPONENT
var appComponentCreateCmd = &cobra.Command{
	Use: "create",
	Short: "Create a new appcomponent.",
	Example: "lvl app component create -n myComponentName -c docker -ctype mysql",
	Run: func(cmd *cobra.Command, args []string) {
		
	},
}

//------------------------------------------------- APP COMPONENTS HELPERS (CATEGORY )-------------------------------------------------

// ---- GET COMPONENT CATEGORIES 
var appComponentCategoryGetCmd = &cobra.Command{
	Use: "categories",
	Short: "shows a list of all current appcomponent categories.",
	Example: "lvl app component categories",
	Run: func(cmd *cobra.Command, args []string) {

		// current possible categories for appcomponents 
		categories := []string{"web-apps", "databases", "extensions"}

		// type to convert string into category type
		var AppcomponentCategories struct{
			Data []types.AppcomponentCategory 
		}


		for _, category := range categories{
			cat := types.AppcomponentCategory{Name: category}
			AppcomponentCategories.Data = append(AppcomponentCategories.Data, cat)
		}
		// display output in readable table
		outputFormatTable(AppcomponentCategories.Data, []string{"CATEGORY"}, []string{"Name"})
	},
}

// ---- GET LIST OF APPCOMPONENT TYPES
var appComponentTypeCmd = &cobra.Command{
	Use: "types",
	Short: "Shows a list of all current componenttypes.",
	Example: "lvl app component types",
	Run: func(cmd *cobra.Command, args []string) {
	
		types := Level27Client.AppComponenttypesGet()

		log.Print(types)
	},
}




	
//-------------------------------------------------  APP SSL CERTIFICATES (GET / CREATE) -------------------------------------------------
// ---- SSL COMMAND
var appCertificateCmd = &cobra.Command{
	Use: "ssl",
	Short: "Commands for managing ssl certificates.",
	Example: "lvl app ssl get",
}


// ---- GET LIST OF SSL CERTIFICATES
var appCertificateGetCmd = &cobra.Command{
	Use: "get",
	Short: "Show list of all available ssl certificates.",
	Example: "lvl app ssl get",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Check for valid AppId type
		appId := checkSingleIntID(args[0], "app")

		certificates := Level27Client.AppCertificateGet(appId)
		// Display output in readable table
		outputFormatTableFuncs(certificates, []string{"NAME", "TYPE", "STATUS", "EXPIRING DATE"}, 
		[]interface{}{"Name", "SslType", "Status", func(s types.SslCertificates) string { return utils.FormatUnixTime(s.DtExpires)}})

	},
}

// ---- ADD CERTIFICATE
// vars to hold properties set by user.
var appAddSslName, appAddSslType, appAddSslKey, appAddSslCrt, appAddSslCabundle string
var appAddSslAutoUrl []string
var appAddSslAutoUrlLink, appAddSslForce bool
var appCertificateAddCmd = &cobra.Command{
	Use: "add",
	Short: "Add new ssl certificate to an app.",
	Example: "lvl app ssl add 2077 -n mySslCertificateName",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Check for valid AppId type
		appId := checkSingleIntID(args[0], "app")
		
		fmt.Print(appId)
	},
}