package types

// main structure of an app
type App struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Status         string `json:"status"`
	StatusCategory string `json:"statusCategory"`
	Organisation   struct {
		ID       int    `json:"id"`
		Name     string `json:"name"`
		Reseller struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"reseller"`
	} `json:"organisation"`
	DtExpires     int    `json:"dtExpires"`
	BillingStatus string `json:"billingStatus"`
	Components    []struct {
		ID               int    `json:"id"`
		Name             string `json:"name"`
		Category         string `json:"category"`
		AppComponentType string `json:"appcomponenttype"`
	} `json:"components"`
	CountTeams int `json:"countTeams"`
	Teams      []struct {
		ID             int    `json:"id"`
		Name           string `json:"name"`
		AdminOnly      bool   `json:"adminOnly"`
		OrganisationID int    `json:"organisationId"`
	} `json:"teams"`
	ExternalInfo string `json:"externalInfo"`
}

//type to create an app (post request)
type AppPostRequest struct {
	Name         string `json:"name"`
	Organisation int    `json:"organisation"`
	AutoTeams    []int  `json:"autoTeams"`
	ExternalInfo string `json:"externalInfo"`
}

//type to update an app (put request)
type AppPutRequest struct {
	Name         string   `json:"name"`
	Organisation int      `json:"organisation"`
	AutoTeams    []string `json:"autoTeams"`
}

// type needed to do an action on a system
type AppActionRequest struct {
	Type string `json:"type"`
}

type AppSslCertificate struct {
	ID                 int         `json:"id"`
	Name               string      `json:"name"`
	SslType            string      `json:"sslType"`
	SslKey             string      `json:"sslKey"`
	NewSslKey          string      `json:"newSslKey"`
	SslCrt             string      `json:"sslCrt"`
	SslCabundle        string      `json:"sslCabundle"`
	AutoURLLink        bool        `json:"autoUrlLink"`
	SslForce           bool        `json:"sslForce"`
	SslStatus          string      `json:"sslStatus"`
	Status             string      `json:"status"`
	ReminderStatus     string      `json:"reminderStatus"`
	DtExpires          string      `json:"dtExpires"`
	ValidationParams   interface{} `json:"validationParams"`
	Source             interface{} `json:"source"`
	SslCertificateUrls []struct {
		ID                int         `json:"id"`
		Content           string      `json:"content"`
		SslStatus         string      `json:"sslStatus"`
		ErrorMsg          interface{} `json:"errorMsg"`
		SslStatusCategory string      `json:"sslStatusCategory"`
		ValidationType    string      `json:"validationType"`
	} `json:"sslCertificateUrls"`
	BillableitemDetail interface{} `json:"billableitemDetail"`
	StatusCategory     string      `json:"statusCategory"`
	SslStatusCategory  string      `json:"sslStatusCategory"`
	Urls               []struct {
		ID             int    `json:"id"`
		Content        string `json:"content"`
		Status         string `json:"status"`
		StatusCategory string `json:"statusCategory"`
	} `json:"urls"`
	MatchingUrls []string `json:"matchingUrls"`
}

type AppSslCertificateCreate struct {
	Name                   string `json:"name"`
	SslType                string `json:"sslType"`
	AutoSslCertificateUrls string `json:"autoSslCertificateUrls"`
	AutoUrlLink            bool   `json:"autoUrlLink"`
	SslForce               bool   `json:"sslForce"`
}

type AppSslCertificateCreateOwn struct {
	AppSslCertificateCreate
	SslKey      string `json:"sslKey"`
	SslCrt      string `json:"sslCrt"`
	SslCabundle string `json:"sslCabundle"`
}

type AppSslCertificatePut struct {
	Name    string `json:"name"`
	SslType string `json:"sslType"`
}

type AppSslcertificateKey struct {
	SslKey string `json:"sslKey"`
}

//type appcomponent
type AppComponent struct {
	App struct {
		ID             int64  `json:"id"`
		Status         string `json:"status"`
		Name           string `json:"name"`
		StatusCategory string `json:"statusCategory"`
	} `json:"app"`
	AppcomponentparameterDescriptions interface{} `json:"appcomponentparameterDescriptions"`
	Appcomponentparameters            interface{} `json:"appcomponentparameters"`
	Appcomponenttype                  string      `json:"appcomponenttype"`
	BillableitemDetailID              int64       `json:"billableitemDetailId"`
	Category                          string      `json:"category"`
	ID                                int64       `json:"id"`
	Name                              string      `json:"name"`
	Organisation                      struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	} `json:"organisation"`
	Provider struct {
		ID   interface{} `json:"id"`
		Name interface{} `json:"name"`
	} `json:"provider"`
	SelectedSystem interface{} `json:"selectedSystem"`
	Status         string      `json:"status"`
	Systemgroup    interface{} `json:"systemgroup"`
	Systems        []struct {
		Cookbooks []interface{} `json:"cookbooks"`
		Fqdn      string        `josn:"fqdn"`
		ID        int64         `json:"id"`
		Name      string        `json:"name"`
	} `json:"systems"`
}

// type appcomponent category
type AppcomponentCategory struct {
	Name string
}

// type appcomponenttype
type Appcomponenttype map[string]AppcomponenttypeServicetype

type AppcomponenttypeServicetype struct {
	Servicetype struct {
		Name                    string        `json:"name"`
		Cookbook                string        `json:"cookbook"`
		DisplayName             string        `json:"displayName"`
		Description             string        `json:"description"`
		URLPossible             bool          `json:"urlPossible"`
		RestorePossible         bool          `json:"restorePossible"`
		MigrationPossible       bool          `json:"migrationPossible"`
		SelectingSystemPossible bool          `json:"selectingSystemPossible"`
		DisabledOnProduction    bool          `json:"disabledOnProduction"`
		InvisibleOnProduction   bool          `json:"invisibleOnProduction"`
		Runlist                 string        `json:"runlist"`
		AllowedActions          []interface{} `json:"allowedActions"`
		Category                string        `json:"category"`
		Parameters              []struct {
			Name         string      `json:"name"`
			DisplayName  string      `json:"displayName"`
			Description  string      `json:"description"`
			Type         string      `json:"type"`
			DefaultValue interface{} `json:"defaultValue"`
			Readonly     bool        `json:"readonly"`
			DisableEdit  bool        `json:"disableEdit"`
			Required     bool        `json:"required"`
			Category     string      `json:"category"`
		} `json:"parameters"`
	} `json:"servicetype"`
}

// type request to add a sslCertificate to an app.
// this type is specificly used when ssl certificate of type "own" is chosen
type AppSslCertificateTypeOwnRequest struct {
	Name                   string `json:"name"`
	SslType                string `json:"sslType"`
	AutoSslCertificateUrls string `json:"autoSslCertificateUrls"`
	SslKey                 string `json:"sslKey"`
	SslCrt                 string `json:"sslCrt"`
	SslCabundle            string `json:"sslCabundle"`
	AutoUrlLink            bool   `json:"autoUrlLink"`
	SslForce               bool   `json:"sslForce"`
}

// type request to add a sslCertificate to an app.
type AppSslCertificateRequest struct {
	Name                   string `json:"name"`
	SslType                string `json:"sslType"`
	AutoSslCertificateUrls string `json:"autoSslCertificateUrls"`
	AutoUrlLink            bool   `json:"autoUrlLink"`
	SslForce               bool   `json:"sslForce"`
}

// request an action on a ssl certificate from an app
type AppSslCertificateActionRequest struct {
	Type string `json:"type"`
}

// Restore type for an app
type AppComponentRestore struct {
	ID           int         `json:"id"`
	Filename     string      `json:"filename"`
	Size         interface{} `json:"size"`
	DtExpires    interface{} `json:"dtExpires"`
	Status       string      `json:"status"`
	Appcomponent struct {
		ID                     int    `json:"id"`
		Name                   string `json:"name"`
		Appcomponenttype       string `json:"appcomponenttype"`
		Appcomponentparameters struct {
			Username string `json:"username"`
			Pass     string `json:"pass"`
		} `json:"appcomponentparameters"`
		Status string `json:"status"`
		App    struct {
			ID int `json:"id"`
		} `json:"app"`
	} `json:"appcomponent"`
	AvailableBackup struct {
		ID           int    `json:"id"`
		Date         string `json:"date"`
		VolumeUID    string `json:"volumeUid"`
		StorageUID   string `json:"storageUid"`
		Status       int    `json:"status"`
		SnapshotName string `json:"snapshotName"`
		System       struct {
			ID           int         `json:"id"`
			Fqdn         string      `json:"fqdn"`
			CustomerFqdn interface{} `json:"customerFqdn"`
			Name         string      `json:"name"`
		} `json:"system"`
		RestoreSystem struct {
			ID           int         `json:"id"`
			Fqdn         string      `json:"fqdn"`
			CustomerFqdn interface{} `json:"customerFqdn"`
			Name         string      `json:"name"`
		} `json:"restoreSystem"`
	} `json:"availableBackup"`
}

// request type for new restore
type AppComponentRestoreRequest struct {
	Appcomponent    int `json:"appcomponent"`
	AvailableBackup int `json:"availableBackup"`
}

// type availablebackup for an appcomponent
type AppComponentAvailableBackup struct {
	Date          string `json:"date"`
	ID            int    `json:"id"`
	RestoreSystem struct {
		Fqdn string `json:"fqdn"`
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"restoreSystem"`
	SnapshotName   string `json:"snapshotName"`
	Status         string `json:"status"`
	StatusCategory string `json:"statusCategory"`
	StorageUID     string `json:"storageUid"`
	System         struct {
		Fqdn string `json:"fqdn"`
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"system"`
	VolumeUID string `json:"volumeUid"`
}

// type app migration
type AppMigration struct {
	ID                 int         `json:"id"`
	MigrationType      string      `json:"migrationType"`
	DtPlanned          interface{} `json:"dtPlanned"`
	Status             string      `json:"status"`
	ConfirmationStatus int         `json:"confirmationStatus"`
	App                struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"app"`

	MigrationItems []struct {
		ID                   int           `json:"id"`
		Type                 string        `json:"type"`
		Source               string        `json:"source"`
		SourceInformation    string        `json:"sourceInformation"`
		DestinationEntity    string        `json:"destinationEntity"`
		DestinationEntityID  int           `json:"destinationEntityId"`
		Status               string        `json:"status"`
		StatusCategory       string        `json:"statusCategory"`
		Ord                  int           `json:"ord"`
		Sshkey               interface{}   `json:"sshkey"`
		InvestigationResults interface{} `json:"investigationResults"`
		PreparationResults   []interface{} `json:"preparationResults"`
		PresyncResults       []interface{} `json:"presyncResults"`
		MigrationResults     []interface{} `json:"migrationResults"`
		Logs                 interface{} `json:"logs"`
		Appcomponent         struct {
			ID                     int    `json:"id"`
			Name                   string `json:"name"`
			Appcomponenttype       string `json:"appcomponenttype"`
			Appcomponentparameters struct {
				User string `json:"user"`
				Pass string `json:"pass"`
				Host string `json:"host"`
			} `json:"appcomponentparameters"`
			Status         string `json:"status"`
			StatusCategory string `json:"statusCategory"`
		} `json:"appcomponent"`
		SourceExtraData struct {
			Appcomponentparameters struct {
				Pass string `json:"pass"`
				Host string `json:"host"`
				User string `json:"user"`
			} `json:"appcomponentparameters"`
			Status         string `json:"status"`
			StatusCategory string `json:"statusCategory"`
			System         struct {
				ID                    int    `json:"id"`
				Fqdn                  string `json:"fqdn"`
				CustomerFqdn          string `json:"customerFqdn"`
				Name                  string `json:"name"`
				Status                string `json:"status"`
				RunningStatus         string `json:"runningStatus"`
				Osv                   string `json:"osv"`
				StatusCategory        string `json:"statusCategory"`
				RunningStatusCategory string `json:"runningStatusCategory"`
			} `json:"system"`
		} `json:"sourceExtraData"`
		DestinationExtraData struct {
			ID                    int    `json:"id"`
			Name                  string `json:"name"`
			Fqdn                  string `json:"fqdn"`
			CustomerFqdn          string `json:"customerFqdn"`
			Status                string `json:"status"`
			StatusCategory        string `json:"statusCategory"`
			RunningStatus         string `json:"runningStatus"`
			RunningStatusCategory string `json:"runningStatusCategory"`
			Osv                   string `json:"osv"`
		} `json:"destinationExtraData"`
	} `json:"migrationItems"`
}

// request type for new migration
type AppMigrationRequest struct {
	MigrationType      string             `json:"migrationType"`
	DtPlanned          string             `json:"dtPlanned"`
	MigrationItemArray []AppMigrationItem `json:"migrationItemArray"`
}

type AppMigrationItem struct {
	Type                string      `json:"type"`
	Source              string      `json:"source"`
	SourceInfo          int         `json:"sourceInformation"`
	DestinationEntity   string      `json:"destinationEntity"`
	DestinationEntityId int         `json:"destinationEntityId"`
	Ord                 string      `json:"ord"`
	SshKey              interface{} `json:"sshKey"`
}

// type appMigration for update
type AppMigrationUpdate struct {
	MigrationType string `json:"migrationType"`
	DtPlanned     string `json:"dtPlanned"`
}

// used to create migration key value pairs
type AppMigrationItemValue map[string]interface{}