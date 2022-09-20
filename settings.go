package main

type Settings struct {
	TrueNasPingUrl   string `envconfig:"TRUENAS_PING_URL"`
	IpmiLoginUrl     string `envconfig:"IPMI_LOGIN_URL"`
	IpmiResetUrl     string `envconfig:"IPMI_RESET_URL"`
	IpmiLoginPayload string `envconfig:"IPMI_LOGIN_PAYLOAD"`
	IpmiResetPayload string `envconfig:"IPMI_RESET_PAYLOAD"`
	IpmiUser         string `envconfig:"IPMI_USER"`
	IpmiPassword     string `envconfig:"IPMI_PASSWORD"`
	SidCookie        string
}
