package types

import "encoding/json"

type System struct {
	Id                    int    `json:"id"`
	Uid                   string `json:"uid"`
	Hostname              string `json:"hostname"`
	Fqdn                  string `json:"fqdn"`
	Name                  string `json:"name"`
	Type                  string `json:"type"`
	Status                string `json:"status"`
	StatusCategory        string `json:"statusCategory"`
	RunningStatus         string `json:"runningStatus"`
	RunningStatusCategory string `json:"runningStatusCategory"`
	Cpu                   int    `json:"cpu"`
	Memory                int    `json:"memory"`
	Disk                  string `json:"disk"`
	MonitoringEnabled     bool   `json:"monitoringEnabled"`
	ManagementType        string `json:"managementType"`
	Organisation          struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"organisation"`
	SystemImage struct {
		Id          int    `json:"id"`
		Name        string `json:"name"`
		ExternalId  string `json:"externalId"`
		OsId        int    `json:"osId"`
		OsName      string `json:"osName"`
		OsType      string `json:"osType"`
		OsVersion   string `json:"osVersion"`
		OsVersionId int    `json:"osVersionId"`
	} `json:"systemimage"`
	OperatingSystemVersion struct {
		Id        int    `json:"id"`
		OsId      int    `json:"osId"`
		OsName    string `json:"osName"`
		OsType    string `json:"osType"`
		OsVersion string `json:"osVersion"`
	} `json:"operatingsystemVersion"`
	ProvideId                   int         `json:"providerId"`
	Provider                    interface{} `json:"provider"`
	ProviderApi                 string      `json:"providerApi"`
	SystemProviderConfiguration struct {
		Id          int    `json:"id"`
		ExternalId  string `json:"externalId"`
		Name        string `json:"name"`
		Description string `json:"description"`
	} `json:"systemproviderConfiguration"`
	Region string `json:"region"`
	Zone   struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	} `json:"zone"`
	Networks         []interface{} `json:"networks"`
	PublicNetworking bool          `json:"publicNetworking"`
	StatsSummary     struct {
		DiskSpace struct {
			Unit  string      `json:"unit"`
			Value interface{} `json:"value"`
			Max   interface{} `json:"max"`
		} `json:"diskspace"`
		Memory struct {
			Unit  string      `json:"unit"`
			Value interface{} `json:"value"`
			Max   interface{} `json:"max"`
		} `json:"Memory"`
		Cpu struct {
			Unit  string      `json:"unit"`
			Value interface{} `json:"value"`
			Max   interface{} `json:"max"`
		} `json:"cpu"`
	} `json:"statsSummary"`
	DtExpires     int    `json:"dtExpires"`
	BillingStatus string `json:"billingStatus"`
	ExternalInfo  string `json:"externalInfo"`
	Remarks       string `json:"remarks"`
}

// data needed for POST request (create system)
type SystemPost struct {
	Name                        string `json:"name"`
	CustomerFqdn                string `json:"customerFqdn"`
	Remarks                     string `json:"remarks"`
	Disk                        *int   `json:"disk"`
	Cpu                         *int   `json:"cpu"`
	Memory                      *int   `json:"memory"`
	MamanagementType            string `json:"managementType"`
	PublicNetworking            bool   `json:"publicNetworking"`
	SystemImage                 int    `json:"systemimage"`
	Organisation                int    `json:"organisation"`
	SystemProviderConfiguration int    `json:"systemproviderConfiguration"`
	Zone                        int    `json:"zone"`
	// InstallSecurityUpdates      *int           `json:"installSecurityUpdates"`
	AutoTeams              string        `json:"autoTeams"`
	ExternalInfo           string        `json:"externalInfo"`
	OperatingSystemVersion *int          `json:"operatingsystemVersion"`
	ParentSystem           *int          `json:"parentsystem"`
	Type                   string        `json:"type"`
	AutoNetworks           []interface{} `json:"autoNetworks"`
}

// ----------------------------------- CHECKS ----------------------------------
type SystemCheck struct {
	SystemCheckData
	CheckParameters interface{} `json:"checkparameters"`
}

type SystemCheckData struct {
	Id                   int    `json:"id"`
	CheckType            string `json:"checktype"`
	ChecktypeLocation    string `json:"checktypeLocation"`
	Status               string `json:"status"`
	StatusInformation    string `json:"statusInformation"`
	DtLastMonitorEnabled int    `json:"dtLastMonitoringEnabled"`
	DtLastStatusChanged  int    `json:"dtLastStatusChange"`
	DtNextCheck          int    `json:"dtNextCheck"`
	DtLastCheck          int    `json:"dtLastCheck"`
	CheckParameters      struct {
		Port struct {
			Value   string `json:"value"`
			Default bool   `json:"default"`
		} `json:"port"`
		W struct {
			Value   string `json:"value"`
			Default bool   `json:"default"`
		} `json:"w"`
		C struct {
			Value   string `json:"value"`
			Default bool   `json:"default"`
		} `json:"c"`
		H struct {
			Value   string `json:"value"`
			Default bool   `json:"default"`
		} `json:"H"`
	} `json:"checkparameters"`

	CheckParametersDescriptions struct {
		W    string `json:"w"`
		C    string `json:"c"`
		H    string `json:"H"`
		Port string `json:"port"`
	} `json:"checkparameterDescriptions"`
}

func (s *SystemCheck) UnmarshalJSON(data []byte) error {
	err := json.Unmarshal(data, &s.SystemCheckData)
	if err != nil {
		return err
	}




	switch s.CheckParameters {
	case "domain":
		var dat struct {
			Entity Domain `json:"entity"`
		}
		err = json.Unmarshal(data, &dat)
		// n.Entity = dat.Entity
	}

	return err
}

// ---- Check create request
type SystemCheckRequest struct {
	Checktype string `json:"checktype"`
}

// ----------------------------------- COOKBOOKS ----------------------------------

type Cookbook struct {
	Id                             int         `json:"id"`
	CookbookType                   string      `json:"cookbooktype"`
	CookbookParameters             interface{} `json:"cookbookparameters"`
	CookbookParametersDescriptions interface{} `json:"cookbookparameterDescriptions"`
	PreviousCookbookParameters     interface{} `json:"previousCookbookparameters"`
	Status                         string      `json:"status"`
}
