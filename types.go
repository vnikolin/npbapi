package npbapi

//VersionInfo
type VersionInfo struct {
	ImageFormatVersion string  `json:"imageFormatVersion"`
	Uptime             float64 `json:"uptime"`
	ModelName          string  `json:"modelName"`
	InternalVersion    string  `json:"internalVersion"`
	MemTotal           int     `json:"memTotal"`
	MfgName            string  `json:"mfgName"`
	SerialNumber       string  `json:"serialNumber"`
	SystemMacAddress   string  `json:"systemMacAddress"`
	BootupTimestamp    float64 `json:"bootupTimestamp"`
	MemFree            int     `json:"memFree"`
	Version            string  `json:"version"`
	ConfigMacAddress   string  `json:"configMacAddress"`
	IsIntlVersion      bool    `json:"isIntlVersion"`
	InternalBuildId    string  `json:"internalBuildId"`
	HardwareRevision   string  `json:"hardwareRevision"`
	HwMacAddress       string  `json:"hwMacAddress"`
	Architecture       string  `json:"architecture"`
}

//HostInfo
type HostInfo struct {
	Fqdn     string `json:"fqdn"`
	HostName string `json:"hostname"`
}

//NtpInfo
type NtpInfo struct {
	Status            string `json:"status"`
	PollingInterfal   int    `json:"pollingInterval"`
	Stratum           int    `json:"stratum"`
	MaxEstimatedError int    `json:"maxEstimatedError"`
	Server            string `json:"server"`
}

//RawPowerInfo
type RawPowerInfo struct {
	PowerSupplies struct {
		Instance1 struct {
			State string `json:"state"`
		} `json:"1"`
		Instance2 struct {
			State string `json:"state"`
		} `json:"2"`
	} `json:"powerSupplies"`
}

//PowerInfo
type PowerInfo struct {
	PS1 string
	PS2 string
}
