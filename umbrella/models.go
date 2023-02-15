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

type NetworkTunnel struct {
	Id           int64        `json:"id"`
	Uri          string       `json:"uri,omitempty"`
	Name         string       `json:"name,omitempty"`
	Siteoriginid int64        `json:"siteOriginId,omitempty"`
	Client       TunnelClient `json:"client,omitempty"`
	Transport    TunnelTrans  `json:"transport,omitempty"`
	Servicetype  string       `json:"serviceType,omitempty"`
	NetworkCIDRs []string     `json:"networkCIDRs,omitempty"`
	Meta         Meta         `json:"meta,omitempty"`
	Createdat    string       `json:"cretedAt,omitempty"`
	Modifiedat   string       `json:"modifiedAt,omitempty"`
}

type NetworkTunnelCreate struct {
	Id             int64       `json:"id"`
	Uri            string      `json:"uri,omitempty"`
	Name           string      `json:"name,omitempty"`
	Siteoriginid   int64       `json:"siteOriginId,omitempty"`
	Devicetype     string      `json:"devicetype,omitempty"`
	Transport      TunnelTrans `json:"transport,omitempty"`
	Servicetype    string      `json:"serviceType,omitempty"`
	NetworkCIDRs   []string    `json:"networkCIDRs,omitempty"`
	Meta           Meta        `json:"meta,omitempty"`
	Createdat      string      `json:"cretedAt,omitempty"`
	Modifiedat     string      `json:"modifiedAt,omitempty"`
	Authentication TunnelAuth  `json:"authentication,omitempty"`
}

type TunnelClient struct {
	Devicetype     string     `json:"deviceType,omitempty"`
	Authentication TunnelAuth `json:"authentication,omitempty"`
}

type TunnelAuth struct {
	Type       string           `json:"type,omitempty"`
	Parameters TunnelAuthParams `json:"parameters,omitempty"`
}

type TunnelAuthParams struct {
	Id         string `json:"id,omitempty"`
	Modifiedat string `json:"modifiedAt,omitempty"`
	Idprefix   string `json:"idPrefix,omitempty"`
	Secret     string `json:"secret,omitempty"`
}

type TunnelTrans struct {
	Protocol string `json:"protocol,omitempty"`
}

type Meta struct{}

type Response struct {
	Message string `json:"message,omitempty"`
}
