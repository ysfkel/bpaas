package config

type ClusterDetail struct {
	CertificateAuthorityData string `mapstructure:"certificate-authority-data"`
	Server                   string `json:"server"`
}
type Cluster struct {
	Name    string        `json:"name"`
	Cluster ClusterDetail `json:"cluster"`
}

type UserDetail struct {
	ClientKeyData         string `mapstructure:"client-key-data"`
	ClientCertificateData string `mapstructure:"client-certificate-data"`
}

type User struct {
	Name string     `json:"name"`
	User UserDetail `json:"user"`
}

type Configuration struct {
	ApiVersion string     `json:"apiVersion"`
	Clusters   []*Cluster `json:"clusters"`
	Users      []*User    `json:"users"`
}
