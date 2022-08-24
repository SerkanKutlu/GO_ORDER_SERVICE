package config

type RemoteServicesConfig struct {
	CustomerService HttpClient `yaml:"customerService"`
}
type HttpClient struct {
	Name    string `yaml:"name"`
	BaseUrl string `yaml:"baseUrl"`
}
