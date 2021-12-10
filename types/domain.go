package types

import (
	"encoding/json"
)

type StructDomain struct {
	ID                    int    `json:"id"`
	Name                  string `json:"name"`
	Fullname              string `json:"fullname"`
	TTL                   int    `json:"ttl"`
	EppCode               string `json:"eppCode"`
	Status                string `json:"status"`
	DnssecStatus          string `json:"dnssecStatus"`
	RegistrationIsHandled bool   `json:"registrationIsHandled"`
	Provider              struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		API  string `json:"api"`
	} `json:"provider"`
	DNSIsHandled    bool   `json:"dnsIsHandled"`
	DtRegister      string `json:"dtRegister"`
	Nameserver1     string `json:"nameserver1"`
	Nameserver2     string `json:"nameserver2"`
	Nameserver3     string `json:"nameserver3"`
	Nameserver4     string `json:"nameserver4"`
	NameserverIP1   string `json:"nameserverIp1"`
	NameserverIP2   string `json:"nameserverIp2"`
	NameserverIP3   string `json:"nameserverIp3"`
	NameserverIP4   string `json:"nameserverIp4"`
	NameserverIpv61 string `json:"nameserverIpv61"`
	NameserverIpv62 string `json:"nameserverIpv62"`
	NameserverIpv63 string `json:"nameserverIpv63"`
	NameserverIpv64 string `json:"nameserverIpv64"`
	Organisation    struct {
		ID       int    `json:"id"`
		Name     string `json:"name"`
		Reseller int    `json:"reseller"`
	} `json:"organisation"`
	Domaintype struct {
		ID                                  int    `json:"id"`
		Name                                string `json:"name"`
		Extension                           string `json:"extension"`
		RenewPeriod                         int    `json:"renewPeriod"`
		TransferAutoLicensee                bool   `json:"transferAutoLicensee"`
		RequestIncomingTransferCodePossible bool   `json:"requestIncomingTransferCodePossible"`
		RequestOutgoingTransferCodePossible bool   `json:"requestOutgoingTransferCodePossible"`
		LicenseeChangePossible              bool   `json:"licenseeChangePossible"`
		DnssecSupported                     bool   `json:"dnssecSupported"`
	} `json:"domaintype"`
	DomaincontactLicensee struct {
		ID               int    `json:"id,omitempty"`
		FirstName        string `json:"firstName"`
		LastName         string `json:"lastName"`
		Fullname         string `json:"fullname"`
		OrganisationName string `json:"organisationName"`
		Street           string `json:"street"`
		HouseNumber      string `json:"houseNumber"`
		Zip              string `json:"zip"`
		City             string `json:"city"`
		State            string `json:"state"`
		Phone            string `json:"phone"`
		Fax              string `json:"fax"`
		Email            string `json:"email"`
		TaxNumber        string `json:"taxNumber"`
		Status           int    `json:"status"`
		PassportNumber   string `json:"passportNumber"`
		SocialNumber     string `json:"socialNumber"`
		BirthStreet      string `json:"birthStreet"`
		BirthZip         string `json:"birthZip"`
		BirthCity        string `json:"birthCity"`
		BirthDate        string `json:"birthDate"`
		Gender           string `json:"gender"`
		Type             string `json:"type"`
		Country          struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"country"`
	} `json:"domaincontactLicensee"`
	DomaincontactOnsite interface{} `json:"domaincontactOnsite"`
	Mailgroup           struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"mailgroup"`
	ExtraFields   []interface{} `json:"extraFields"`
	HandleMailDNS bool          `json:"handleMailDns"`
	DtExpires     int           `json:"dtExpires"`
	BillingStatus string        `json:"billingStatus"`
	ExternalInfo  string        `json:"externalInfo"`
	Teams         []struct {
		ID             int    `json:"id"`
		Name           string `json:"name"`
		AdminOnly      bool   `json:"adminOnly"`
		OrganisationId int    `json:"organisationId"`
	} `json:"teams"`
	CountTeams int `json:"countTeams"`
}

// Domain represents a single domain
type Domain struct {
	Data StructDomain `json:"domain"`
}

// Domains represents an array of domains
type Domains struct {
	Data []StructDomain `json:"domains"`
}

func (d Domain) String() string {
	return "domain"
}

// DomainProvider represents a single DomainProvider
type DomainProvider struct {
	Providers []struct {
		ID              int    `json:"id"`
		Name            string `json:"name"`
		API             string `json:"api"`
		DNSSecSupported bool   `json:"dnsSecSupported"`
		Domaintypes     []struct {
			ID        int    `json:"id"`
			Extension string `json:"extension"`
		} `json:"domaintypes"`
	} `json:"providers"`
}

// DomainExtension represents a single DomainExtension
type DomainExtension struct {
	ID        int
	Extension string
}

// DomainRequest represents a single DomainRequest
type DomainRequest struct {
	Name                      string `json:"name"`
	NameServer1               string `json:"nameserver1"`
	NameServer2               string `json:"nameserver2"`
	NameServer3               string `json:"nameserver3"`
	NameServer4               string `json:"nameserver4"`
	NameServer1Ip             string `json:"nameserverIp1"`
	NameServer2Ip             string `json:"nameserverIp2"`
	NameServer3Ip             string `json:"nameserverIp3"`
	NameServer4Ip             string `json:"nameserverIp4"`
	NameServer1Ipv6           string `json:"nameserverIpv61"`
	NameServer2Ipv6           string `json:"nameserverIpv62"`
	NameServer3Ipv6           string `json:"nameserverIpv63"`
	NameServer4Ipv6           string `json:"nameserverIpv64"`
	TTL                       int    `json:"ttl"`
	Action                    string `json:"action"`
	EppCode                   string `json:"eppCode"`
	Handledns                 bool   `json:"handleDns"`
	ExtraFields               string `json:"extraFields"`
	Domaintype                int    `json:"domaintype"`
	Domaincontactlicensee     int    `json:"domaincontactLicensee"`
	DomainContactOnSite       string    `json:"domaincontactOnsite"`
	Organisation              int    `json:"organisation"`
	AutoRecordTemplate        string `json:"autorecordTemplate"`
	AutoRecordTemplateReplace bool   `json:"autorecordTemplateReplace"`
	DomainProvider            int    `json:"domainprovider"`
	DtExternalCreated         string `json:"dtExternalCreated"`
	DtExternalExpires         string `json:"dtExternalExpires"`
	ConvertDomainRecords      string `json:"convertDomainrecords"`
	AutoTeams                 string `json:"autoTeams"`
	ExternalInfo              string `json:"ExternalInfo"`
}

func (d DomainRequest) String() string {
	
	s, _ := json.Marshal(d)
	return string(s)

	// s := "{"
	// s += fmt.Sprintf("\"name\": \"%s\",", d.Name)
	// s += fmt.Sprintf("\"nameserver1\": \"%s\",", d.NameServer1)
	// s += fmt.Sprintf("\"nameserver2\": \"%s\",", d.NameServer2)
	// s += fmt.Sprintf("\"nameserver3\": \"%s\",", d.NameServer3)
	// s += fmt.Sprintf("\"nameserver4\": \"%s\",", d.NameServer4)
	// s += fmt.Sprintf("\"nameserverIp1\": \"%s\",", d.NameServer1Ip)
	// s += fmt.Sprintf("\"nameserverIp2\": \"%s\",", d.NameServer2Ip)
	// s += fmt.Sprintf("\"nameserverIp3\": \"%s\",", d.NameServer3Ip)
	// s += fmt.Sprintf("\"nameserverIp4\": \"%s\",", d.NameServer4Ip)
	// s += fmt.Sprintf("\"nameserverIpv61\": \"%s\",", d.NameServer1Ipv6)
	// s += fmt.Sprintf("\"nameserverIpv62\": \"%s\",", d.NameServer2Ipv6)
	// s += fmt.Sprintf("\"nameserverIpv63\": \"%s\",", d.NameServer3Ipv6)
	// s += fmt.Sprintf("\"nameserverIpv64\": \"%s\",", d.NameServer4Ipv6)
	// s += fmt.Sprintf("\"action\": \"%s\",", d.Action)
	// s += fmt.Sprintf("\"ttl\": \"%v\",", d.TTL)
	// s += fmt.Sprintf("\"eppCode\": \"%s\",", d.EppCode)
	// s += fmt.Sprintf("\"handleDns\": \"%t\",", d.Handledns)
	// s += fmt.Sprintf("\"extraFields\": \"%s\",", d.ExtraFields)
	// s += fmt.Sprintf("\"domaintype\": \"%d\",", d.Domaintype)
	// s += fmt.Sprintf("\"domaincontactLicensee\": \"%v\",", d.Domaincontactlicensee)
	// s += fmt.Sprintf("\"domaincontactOnsite\": \"%v\",", d.DomainContactOnSite)
	// s += fmt.Sprintf("\"organisation\": \"%v\",", d.Organisation)
	// s += fmt.Sprintf("\"autoRecordTemplate\": \"%s\",", d.AutoRecordTemplate)
	// s += fmt.Sprintf("\"autoRecordTemplateReplace\": \"%v\",", d.AutoRecordTemplateReplace)
	// s += fmt.Sprintf("\"domainprovider\": \"%v\",", d.DomainProvider)
	// s += fmt.Sprintf("\"dtExternalCreated\": \"%s\",", d.DtExternalCreated)
	// s += fmt.Sprintf("\"dtExternalExpires\": \"%s\",", d.DtExternalExpires)
	// s += fmt.Sprintf("\"convertDomainRecords\": \"%s\",", d.ConvertDomainRecords)
	// s += fmt.Sprintf("\"autoTeams\": \"%s\",", d.AutoTeams)
	// s += fmt.Sprintf("\"externalInfo\": \"%s\",", d.ExternalInfo)

	// s += "}"
	// return s
}

// DomainRecord represents a single Domainrecord
type DomainRecord struct {
	ID                 int    `json:"id"`
	Name               string `json:"name"`
	Content            string `json:"content"`
	Priority           int    `json:"priority"`
	Type               string `json:"type"`
	SystemHasNetworkIP struct {
		ID int `json:"id"`
	} `json:"systemHasNetworkIp"`
	// URL            int `json:"url"`
	// SslCertificate int `json:"sslCertificate"`
	// Mailgroup      int `json:"mailgroup"`
}

// DomainRecordRequest represents a API reqest to Level27
type DomainRecordRequest struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Content  string `json:"content"`
	Priority int    `json:"priority"`
}

// DomainContact is an object to define domain contacts at Level27
type DomainContact struct {
	Domaincontact struct {
		ID               int    `json:"id"`
		FirstName        string `json:"firstName"`
		LastName         string `json:"lastName"`
		OrganisationName string `json:"organisationName"`
		Street           string `json:"street"`
		HouseNumber      string `json:"houseNumber"`
		Zip              string `json:"zip"`
		City             string `json:"city"`
		State            string `json:"state"`
		Phone            string `json:"phone"`
		Fax              string `json:"fax"`
		Email            string `json:"email"`
		TaxNumber        string `json:"taxNumber"`
		PassportNumber   string `json:"passportNumber"`
		SocialNumber     string `json:"socialNumber"`
		BirthStreet      string `json:"birthStreet"`
		BirthZip         string `json:"birthZip"`
		BirthCity        string `json:"birthCity"`
		BirthDate        string `json:"birthDate"`
		Gender           string `json:"gender"`
		Type             string `json:"type"`
		Country          struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"country"`
		Organisation struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"organisation"`
		Fullname string `json:"fullname"`
	} `json:"domaincontact"`
}

// DomainContactRequest is an object to define the request to create or modify a domain contact at Level27
type DomainContactRequest struct {
	Type             string `json:"type"`
	FirstName        string `json:"firstName"`
	LastName         string `json:"lastName"`
	OrganisationName string `json:"organisationName"`
	Street           string `json:"street"`
	HouseNumber      string `json:"houseNumber,omitempty"`
	Zip              string `json:"zip"`
	City             string `json:"city"`
	State            string `json:"state,omitempty"`
	Phone            string `json:"phone"`
	Fax              string `json:"fax,omitempty"`
	Email            string `json:"email"`
	TaxNumber        string `json:"taxNumber"`
	PassportNumber   string `json:"passportNumber,omitempty"`
	SocialNumber     string `json:"socialNumber,omitempty"`
	BirthStreet      string `json:"birthStreet,omitempty"`
	BirthZip         string `json:"birthZip,omitempty"`
	BirthCity        string `json:"birthCity,omitempty"`
	BirthDate        string `json:"birthDate,omitempty"`
	Gender           string `json:"gender,omitempty"`
	Country          string `json:"country"`
	Organisation     string `json:"organisation"`
}

func (d DomainContactRequest) String() string {
	s, _ := json.Marshal(d)
	return string(s)
}
