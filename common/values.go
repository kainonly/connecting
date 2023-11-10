package common

import (
	"github.com/weplanx/go/values"
	"strings"
)

type Values struct {
	Mode      string `env:"MODE" envDefault:"debug"`
	Hostname  string `env:"HOSTNAME"`
	Address   string `env:"ADDRESS"`
	Console   string `env:"CONSOLE,required"`
	Ip        string `env:"IP" envDefault:"X-Forwarded-For"`
	XDomain   string `env:"XDOMAIN" envDefault:""`
	Namespace string `env:"NAMESPACE,required"`
	Key       string `env:"KEY,required"`

	Database struct {
		Url   string `env:"URL,required"`
		Name  string `env:"NAME,required"`
		Redis string `env:"REDIS,required"`
	} `envPrefix:"DATABASE_"`

	Nats struct {
		Hosts []string `env:"HOSTS,required" envSeparator:","`
		Nkey  string   `env:"NKEY,required"`
	} `envPrefix:"NATS_"`

	Otlp struct {
		Endpoint string `env:"ENDPOINT,required"`
		Token    string `env:"TOKEN,required"`
	} `envPrefix:"OTLP_"`

	*Extra
}

type Extra struct {
	IpAddress             string `yaml:"ip_address"`
	IpSecretId            string `yaml:"ip_secret_id"`
	IpSecretKey           string `yaml:"ip_secret_key" secret:"*"`
	Ipv6Address           string `yaml:"ipv6_address" json:"Ipv6Address"`
	Ipv6SecretId          string `yaml:"ipv6_secret_id" json:"Ipv6SecretId"`
	Ipv6SecretKey         string `yaml:"ipv6_secret_key" secret:"*" json:"Ipv6SecretKey"`
	SmsSecretId           string `yaml:"sms_secret_id"`
	SmsSecretKey          string `yaml:"sms_secret_key" secret:"*"`
	SmsSign               string `yaml:"sms_sign"`
	SmsAppId              string `yaml:"sms_app_id"`
	SmsRegion             string `yaml:"sms_region"`
	SmsPhoneBind          string `yaml:"sms_phone_bind"`
	SmsLoginVerify        string `yaml:"sms_login_verify"`
	*values.DynamicValues `yaml:"dynamic_values"`
}

func (x Values) IsRelease() bool {
	return x.Mode == "release"
}

func (x Values) NameX(sep string, v ...string) string {
	elems := []string{x.Namespace}
	elems = append(elems, v...)
	return strings.Join(elems, sep)
}