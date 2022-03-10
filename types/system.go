package types

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
	ProvideId                   int                            `json:"providerId"`
	Provider                    interface{}                    `json:"provider"`
	ProviderApi                 string                         `json:"providerApi"`
	SystemProviderConfiguration SystemProviderConfigurationRef `json:"systemproviderConfiguration"`
	Region                      string                         `json:"region"`
	Zone                        struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	} `json:"zone"`
	Networks         []SystemNetwork `json:"networks"`
	PublicNetworking bool            `json:"publicNetworking"`
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
	Groups        []struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	} `json:"groups"`
	Jobs         []Job `json:"jobs"`
	ParentSystem *struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	} `json:"parentsystem"`
	InstallSecurityUpdates int `json:"installSecurityUpdates"`
	LimitRiops             int `json:"limitRiops"`
	LimitWiops             int `json:"limitWiops"`
	BootVolume             struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	} `json:"bootVolume"`
	Cookbooks             []SystemCookbook `json:"cookbooks"`
	Preferredparentsystem string           `json:"preferredparentsystem"`
}

type DescribeSystem struct {
	System
	SshKeys                      []SystemSshkey     `json:"sshKeys"`
	InstallSecurityUpdatesString string             `json:"installSecurityUpdatesString"`
	HasNetworks                  []SystemHasNetwork `json:"hasNetworks"`
	Volumes                      []SystemVolume     `json:"volumes"`
}
type SystemVolume struct {
	ID           int         `json:"id"`
	Name         string      `json:"name"`
	Status       string      `json:"status"`
	Space        int         `json:"space"`
	UID          string      `json:"uid"`
	Remarks      interface{} `json:"remarks"`
	AutoResize   bool        `json:"autoResize"`
	DeviceName   string      `json:"deviceName"`
	Organisation struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"organisation"`
	System struct {
		ID   int    `json:"id"`
		Fqdn string `json:"fqdn"`
		Name string `json:"name"`
	} `json:"system"`
	Volumegroup struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"volumegroup"`
	StatusCategory string `json:"statusCategory"`
}

type SshKey struct {
	Id           int             `json:"id"`
	Description  string          `json:"description"`
	Content      string          `json:"content"`
	Status       string          `json:"status"`
	Fingerprint  string          `json:"fingerprint"`
	Organisation OrganisationRef `json:"organisation"`
}

type SystemSshkey struct {
	ID           int             `json:"id"`
	Description  string          `json:"description"`
	Fingerprint  string          `json:"fingerprint"`
	Organisation OrganisationRef `json:"organisation"`
	User         struct {
		ID             int    `json:"id"`
		FirstName      string `json:"firstName"`
		LastName       string `json:"lastName"`
		Status         string `json:"status"`
		StatusCategory string `json:"statusCategory"`
	} `json:"user"`
	ShsID             int    `json:"shsId"`
	ShsStatusCategory string `json:"shsStatusCategory"`
	ShsStatus         string `json:"shsStatus"`
}

type SystemNetwork struct {
	ID           int    `json:"id"`
	Mac          string `json:"mac"`
	NetworkID    int    `json:"networkId"`
	Name         string `json:"name"`
	UID          string `json:"uid"`
	NetIpv4      string `json:"netIpv4"`
	NetGatewayv4 string `json:"netGatewayv4"`
	NetMaskv4    int    `json:"netMaskv4"`
	NetIpv6      string `json:"netIpv6"`
	NetGatewayv6 string `json:"netGatewayv6"`
	NetMaskv6    int    `json:"netMaskv6"`
	NetPublic    bool   `json:"netPublic"`
	NetCustomer  bool   `json:"netCustomer"`
	NetInternal  bool   `json:"netInternal"`
	Vlan         int    `json:"vlan"`
	Ips          []struct {
		ID         int    `json:"id"`
		PublicIpv4 string `json:"publicIpv4"`
		Ipv4       string `json:"ipv4"`
		PublicIpv6 string `json:"publicIpv6"`
		Ipv6       string `json:"ipv6"`
		Hostname   string `json:"hostname"`
	} `json:"ips"`
	Destinationv4 []string `json:"destinationv4"`
	Destinationv6 []string `json:"destinationv6"`
	NetslotNumber int      `json:"netslotNumber"`
}

type SystemHasNetwork struct {
	ID             int         `json:"id"`
	Mac            string      `json:"mac"`
	Status         string      `json:"status"`
	StatusCategory string      `json:"statusCategory"`
	ExternalID     interface{} `json:"externalId"`
	Network        NetworkRef  `json:"network"`
}

type SystemCookbook struct {
	ID                 int    `json:"id"`
	Cookbooktype       string `json:"cookbooktype"`
	Cookbookparameters map[string]struct {
		Value   interface{} `json:"value"`
		Default bool        `json:"default"`
	} `json:"cookbookparameters"`
	CookbookparameterDescriptions map[string]string `json:"cookbookparameterDescriptions"`
	PreviousCookbookparameters    string            `json:"previousCookbookparameters"`
	Status                        string            `json:"status"`
	StatusCategory                string            `json:"statusCategory"`
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

type SystemCheckTypeName map[string]SystemCheckType

type SystemCheckType struct {
	ServiceType struct {
		Name            string `json:"name"`
		DisplayName     string `json:"displayName"`
		Description     string `json:"descriptiom"`
		Location        string `json:"location"`
		AlwaysApply     bool   `json:"alwaysApply"`
		OperatingSystem string `json:"operatingSystem"`
		EntityType      string `json:"entityType"`
		Parameters      []struct {
			Name         string      `json:"name"`
			Description  string      `json:"description"`
			Type         string      `json:"type"`
			DefaultValue interface{} `json:"defaultValue"`
			Mandatory    bool        `json:"mandatory"`
		} `json:"parameters"`
	} `json:"servicetype"`
}

type SystemCheck struct {
	Id                   int    `json:"id"`
	CheckType            string `json:"checktype"`
	ChecktypeLocation    string `json:"checktypeLocation"`
	Status               string `json:"status"`
	StatusInformation    string `json:"statusInformation"`
	DtLastMonitorEnabled int    `json:"dtLastMonitoringEnabled"`
	DtLastStatusChanged  int64  `json:"dtLastStatusChange"`
	DtNextCheck          int    `json:"dtNextCheck"`
	DtLastCheck          int    `json:"dtLastCheck"`
	CheckParameters      interface {
	} `json:"checkparameters"`
	CheckParametersDescriptions interface {
	} `json:"checkparameterDescriptions"`
	Location string `json:"location"`
	System   struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	} `json:"system"`
	Alerts []interface{} `json:"alerts"`
}

// ---- Check create request
type SystemCheckRequest struct {
	Checktype string `json:"checktype"`
}

// ---- check create request for http type
type SystemCheckRequestHttp struct {
	Checktype string `json:"checktype"`
	Port      int    `json:"port"`
	Hostname  string `json:"hostname"`
	Url       string `json:"url"`
	Content   string `json:"content"`
}

// ----------------------------------- COOKBOOKS ----------------------------------
// parameteroptions
type CookbookParameterOption struct {
	Name                    string      `json:"name"`
	Exclusive               bool        `json:"exclusive"`
	Value                   interface{} `json:"value"`
	OperatingSystemVersions []struct {
		Name    string `json:"name"`
		Default bool   `json:"default"`
	} `json:"operatingsystem_versions"`
}


// Cookbooktype (used to see all current valid cookbooktypes)
type CookbookTypeName map[string]CookbookType
type CookbookType struct {
	CookbookType struct {
		Name        string `json:"name"`
		DisplayName string `json:"displayName"`
		Description string `json:"description"`
		Parameters  []struct {
			Name         string      `json:"name"`
			Description  string      `json:"description"`
			Type         string      `json:"type"`
			DefaultValue interface{} `json:"defaultValue"`
		} `json:"parameters"`
		ParameterOptions interface{} `json:"parameterOptions"`
	} `json:"cookbooktype"`
}

type Cookbook struct {
	Id                             int         `json:"id"`
	CookbookType                   string      `json:"cookbooktype"`
	CookbookParameters             interface{} `json:"cookbookparameters"`
	CookbookParametersDescriptions interface{} `json:"cookbookparameterDescriptions"`
	PreviousCookbookParameters     interface{} `json:"previousCookbookparameters"`
	Status                         string      `json:"status"`
	System                         struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	} `json:"system"`
}

// // Add cookbook to a system request
// type CookbookAdd struct {
// 	Cookbooktype string `json:"cookbooktype"`
// }

type SystemProviderConfigurationRef struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	ExternalID  string `json:"externalId"`
	Description string `json:"description"`
}

type SystemProviderConfiguration struct {
	SystemProviderConfigurationRef
	MinCPU         int    `json:"minCpu"`
	MaxCPU         int    `json:"maxCpu"`
	MinMemory      string `json:"minMemory"`
	MaxMemory      string `json:"maxMemory"`
	MinDisk        int    `json:"minDisk"`
	MaxDisk        int    `json:"maxDisk"`
	Status         int    `json:"status"`
	Systemprovider struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"systemprovider"`
}

type SystemPut struct {
	Id                          int    `json:"id"`
	Name                        string `json:"name"`
	Type                        string `json:"type"`
	Cpu                         int    `json:"cpu"`
	Memory                      int    `json:"memory"`
	Disk                        string `json:"disk"`
	ManagementType              string `json:"managementType"`
	Organisation                int    `json:"organisation"`
	SystemImage                 int    `json:"systemimage"`
	OperatingsystemVersion      int    `json:"operatingsystemVersion"`
	SystemProviderConfiguration int    `json:"systemproviderConfiguration"`
	Zone                        int    `json:"zone"`
	PublicNetworking            bool   `json:"publicNetworking"`
	Preferredparentsystem       string `json:"preferredparentsystem"`
	Remarks                     string `json:"remarks"`
	InstallSecurityUpdates      int    `json:"installSecurityUpdates"`
	LimitRiops                  int    `json:"limitRiops"`
	LimitWiops                  int    `json:"limitWiops"`
}
