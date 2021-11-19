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
