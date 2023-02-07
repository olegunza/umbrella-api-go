package umbrella

type Network struct {
	Name         string `json:"name"`
	Originid     int64  `json:"originId,omitempty"`
	Isdynamic    bool   `json:"isDynamic"`
	Isverified   bool   `json:"isVerified,omitempty"`
	Prefixlength int16  `json:"prefixLength"`
	Ipaddress    string `json:"ipAddress"`
	Status       string `json:"status"`
	Createdat    string `json:"createdAt,omitempty"`
}

type Internalnetwork struct {
	Name         string `json:"name"`
	Ipaddress    string `json:"ipAddress"`
	Prefixlength int16  `json:"prefixLength"`
	Sitename     string `json:"siteName,omitempty"`
	Siteid       int16  `json:"siteId,omitempty"`
	Networkname  string `json:"networkName,omitempty"`
	Networkid    int64  `json:"networkId,omitempty"`
	Tunnelname   string `json:"tunnelName,omitempty"`
	Tunnelid     int64  `json:"tunnelId,omitempty"`
	Createdat    string `json:"createdAt,omitempty"`
	Modifiedat   string `json:"modifiedAt,omitempty"`
}

type Site struct {
	Originid             int64  `json:"originId,omitempty"`
	Name                 string `json:"name"`
	Siteid               int    `json:"siteId,omitempty"`
	Isdefault            bool   `json:"isDefault,omitempty"`
	Type                 string `json:"type,omitempty"`
	Internalnetworkcount int32  `json:"internalNetworkCount,omitempty"`
	Vacount              int32  `json:"vaCount,omitempty"`
	Createdat            string `json:"createdAt,omitempty"`
	Modifiedat           string `json:"modifiedAt,omitempty"`
}
