package mirakurun

// ServicesService ...
type ServicesService service

// Service represents a Mirakurun service.
type Service struct {
	ID                 int     `json:"id"`
	ServiceID          int     `json:"serviceId"`
	NetworkID          int     `json:"networkId"`
	Name               string  `json:"name"`
	LogoID             int     `json:"logoId,omitempty"`
	HasLogoData        bool    `json:"hasLogoData,omitempty"`
	RemoteControlKeyID int     `json:"remoteControlKeyId,omitempty"`
	Channel            Channel `json:"channel,omitempty"`
}
