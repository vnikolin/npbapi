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

//IntInfo
type IntInfo struct {
	InterfaceStatuses struct {
		Ethernet_1_1 struct {
			VlanInformation struct {
				VlanExplanation string `json:"vlanExplanation"`
				InterfaceMode   string `json:"interfaceMode"`
			} `json:"vlanInformation"`
			BandWidth  int    `json:"bandwidth"`
			LinkStatus string `json:"linkStatus"`
		} `json:"Ethernet1/1"`
		Ethernet_2_1 struct {
			VlanInformation struct {
				VlanExplanation string `json:"vlanExplanation"`
				InterfaceMode   string `json:"interfaceMode"`
			} `json:"vlanInformation"`
			BandWidth  int    `json:"bandwidth"`
			LinkStatus string `json:"linkStatus"`
		} `json:"Ethernet2/1"`
		Ethernet_3_1 struct {
			VlanInformation struct {
				VlanExplanation string `json:"vlanExplanation"`
				InterfaceMode   string `json:"interfaceMode"`
			} `json:"vlanInformation"`
			BandWidth  int    `json:"bandwidth"`
			LinkStatus string `json:"linkStatus"`
		} `json:"Ethernet3/1"`
		Ethernet_4_1 struct {
			VlanInformation struct {
				VlanExplanation string `json:"vlanExplanation"`
				InterfaceMode   string `json:"interfaceMode"`
			} `json:"vlanInformation"`
			BandWidth  int    `json:"bandwidth"`
			LinkStatus string `json:"linkStatus"`
		} `json:"Ethernet4/1"`
		Ethernet_5_1 struct {
			VlanInformation struct {
				VlanExplanation string `json:"vlanExplanation"`
				InterfaceMode   string `json:"interfaceMode"`
			} `json:"vlanInformation"`
			BandWidth  int    `json:"bandwidth"`
			LinkStatus string `json:"linkStatus"`
		} `json:"Ethernet5/1"`
		Ethernet_6_1 struct {
			VlanInformation struct {
				VlanExplanation string `json:"vlanExplanation"`
				InterfaceMode   string `json:"interfaceMode"`
			} `json:"vlanInformation"`
			BandWidth  int    `json:"bandwidth"`
			LinkStatus string `json:"linkStatus"`
		} `json:"Ethernet6/1"`
		Ethernet_7_1 struct {
			VlanInformation struct {
				VlanExplanation string `json:"vlanExplanation"`
				InterfaceMode   string `json:"interfaceMode"`
			} `json:"vlanInformation"`
			BandWidth  int    `json:"bandwidth"`
			LinkStatus string `json:"linkStatus"`
		} `json:"Ethernet7/1"`
		Ethernet_8_1 struct {
			VlanInformation struct {
				VlanExplanation string `json:"vlanExplanation"`
				InterfaceMode   string `json:"interfaceMode"`
			} `json:"vlanInformation"`
			BandWidth  int    `json:"bandwidth"`
			LinkStatus string `json:"linkStatus"`
		} `json:"Ethernet8/1"`
		Ethernet_9_1 struct {
			VlanInformation struct {
				VlanExplanation string `json:"vlanExplanation"`
				InterfaceMode   string `json:"interfaceMode"`
			} `json:"vlanInformation"`
			BandWidth  int    `json:"bandwidth"`
			LinkStatus string `json:"linkStatus"`
		} `json:"Ethernet9/1"`
		Ethernet_10_1 struct {
			VlanInformation struct {
				VlanExplanation string `json:"vlanExplanation"`
				InterfaceMode   string `json:"interfaceMode"`
			} `json:"vlanInformation"`
			BandWidth  int    `json:"bandwidth"`
			LinkStatus string `json:"linkStatus"`
		} `json:"Ethernet10/1"`
		Ethernet_11_1 struct {
			VlanInformation struct {
				VlanExplanation string `json:"vlanExplanation"`
			} `json:"vlanInformation"`
			BandWidth  int    `json:"bandwidth"`
			LinkStatus string `json:"linkStatus"`
		} `json:"Ethernet11/1"`
		Ethernet_12_1 struct {
			VlanInformation struct {
				VlanExplanation string `json:"vlanExplanation"`
				InterfaceMode   string `json:"interfaceMode"`
			} `json:"vlanInformation"`
			BandWidth  int    `json:"bandwidth"`
			LinkStatus string `json:"linkStatus"`
		} `json:"Ethernet12/1"`
		Ethernet_13_1 struct {
			VlanInformation struct {
				VlanExplanation string `json:"vlanExplanation"`
				InterfaceMode   string `json:"interfaceMode"`
			} `json:"vlanInformation"`
			BandWidth  int    `json:"bandwidth"`
			LinkStatus string `json:"linkStatus"`
		} `json:"Ethernet13/1"`
		Ethernet_14_1 struct {
			VlanInformation struct {
				VlanExplanation string `json:"vlanExplanation"`
				InterfaceMode   string `json:"interfaceMode"`
			} `json:"vlanInformation"`
			BandWidth  int    `json:"bandwidth"`
			LinkStatus string `json:"linkStatus"`
		} `json:"Ethernet14/1"`
		Ethernet_15_1 struct {
			VlanInformation struct {
				VlanExplanation string `json:"vlanExplanation"`
				InterfaceMode   string `json:"interfaceMode"`
			} `json:"vlanInformation"`
			BandWidth  int    `json:"bandwidth"`
			LinkStatus string `json:"linkStatus"`
		} `json:"Ethernet15/1"`
		Ethernet_16_1 struct {
			VlanInformation struct {
				VlanExplanation string `json:"vlanExplanation"`
				InterfaceMode   string `json:"interfaceMode"`
			} `json:"vlanInformation"`
			BandWidth  int    `json:"bandwidth"`
			LinkStatus string `json:"linkStatus"`
		} `json:"Ethernet16/1"`
		Ethernet_17_1 struct {
			VlanInformation struct {
				VlanExplanation string `json:"vlanExplanation"`
				InterfaceMode   string `json:"interfaceMode"`
			} `json:"vlanInformation"`
			BandWidth  int    `json:"bandwidth"`
			LinkStatus string `json:"linkStatus"`
		} `json:"Ethernet17/1"`
		Ethernet_18_1 struct {
			VlanInformation struct {
				VlanExplanation string `json:"vlanExplanation"`
				InterfaceMode   string `json:"interfaceMode"`
			} `json:"vlanInformation"`
			BandWidth  int    `json:"bandwidth"`
			LinkStatus string `json:"linkStatus"`
		} `json:"Ethernet18/1"`
		Ethernet_19_1 struct {
			VlanInformation struct {
				VlanExplanation string `json:"vlanExplanation"`
				InterfaceMode   string `json:"interfaceMode"`
			} `json:"vlanInformation"`
			BandWidth  int    `json:"bandwidth"`
			LinkStatus string `json:"linkStatus"`
		} `json:"Ethernet19/1"`
		Ethernet_20_1 struct {
			VlanInformation struct {
				VlanExplanation string `json:"vlanExplanation"`
				InterfaceMode   string `json:"interfaceMode"`
			} `json:"vlanInformation"`
			BandWidth  int    `json:"bandwidth"`
			LinkStatus string `json:"linkStatus"`
		} `json:"Ethernet20/1"`
		Ethernet_21_1 struct {
			VlanInformation struct {
				VlanExplanation string `json:"vlanExplanation"`
				InterfaceMode   string `json:"interfaceMode"`
			} `json:"vlanInformation"`
			BandWidth  int    `json:"bandwidth"`
			LinkStatus string `json:"linkStatus"`
		} `json:"Ethernet21/1"`
		Ethernet_22_1 struct {
			VlanInformation struct {
				VlanExplanation string `json:"vlanExplanation"`
				InterfaceMode   string `json:"interfaceMode"`
			} `json:"vlanInformation"`
			BandWidth  int    `json:"bandwidth"`
			LinkStatus string `json:"linkStatus"`
		} `json:"Ethernet22/1"`
		Ethernet_23_1 struct {
			VlanInformation struct {
				VlanExplanation string `json:"vlanExplanation"`
				InterfaceMode   string `json:"interfaceMode"`
			} `json:"vlanInformation"`
			BandWidth  int    `json:"bandwidth"`
			LinkStatus string `json:"linkStatus"`
		} `json:"Ethernet23/1"`
		Ethernet_24_1 struct {
			VlanInformation struct {
				VlanExplanation string `json:"vlanExplanation"`
				InterfaceMode   string `json:"interfaceMode"`
			} `json:"vlanInformation"`
			BandWidth  int    `json:"bandwidth"`
			LinkStatus string `json:"linkStatus"`
		} `json:"Ethernet24/1"`
		Ethernet_25_1 struct {
			VlanInformation struct {
				VlanExplanation string `json:"vlanExplanation"`
				InterfaceMode   string `json:"interfaceMode"`
			} `json:"vlanInformation"`
			BandWidth  int    `json:"bandwidth"`
			LinkStatus string `json:"linkStatus"`
		} `json:"Ethernet25/1"`
		Ethernet_26_1 struct {
			VlanInformation struct {
				VlanExplanation string `json:"vlanExplanation"`
				InterfaceMode   string `json:"interfaceMode"`
			} `json:"vlanInformation"`
			BandWidth  int    `json:"bandwidth"`
			LinkStatus string `json:"linkStatus"`
		} `json:"Ethernet26/1"`
		Ethernet_27_1 struct {
			VlanInformation struct {
				VlanExplanation string `json:"vlanExplanation"`
				InterfaceMode   string `json:"interfaceMode"`
			} `json:"vlanInformation"`
			BandWidth  int    `json:"bandwidth"`
			LinkStatus string `json:"linkStatus"`
		} `json:"Ethernet27/1"`
		Ethernet_28_1 struct {
			VlanInformation struct {
				VlanExplanation string `json:"vlanExplanation"`
				InterfaceMode   string `json:"interfaceMode"`
			} `json:"vlanInformation"`
			BandWidth  int    `json:"bandwidth"`
			LinkStatus string `json:"linkStatus"`
		} `json:"Ethernet28/1"`
		Ethernet_29_1 struct {
			VlanInformation struct {
				VlanExplanation string `json:"vlanExplanation"`
				InterfaceMode   string `json:"interfaceMode"`
			} `json:"vlanInformation"`
			BandWidth  int    `json:"bandwidth"`
			LinkStatus string `json:"linkStatus"`
		} `json:"Ethernet29/1"`
		Ethernet_30_1 struct {
			VlanInformation struct {
				VlanExplanation string `json:"vlanExplanation"`
				InterfaceMode   string `json:"interfaceMode"`
			} `json:"vlanInformation"`
			BandWidth  int    `json:"bandwidth"`
			LinkStatus string `json:"linkStatus"`
		} `json:"Ethernet30/1"`
		Ethernet_31_1 struct {
			VlanInformation struct {
				VlanExplanation string `json:"vlanExplanation"`
				InterfaceMode   string `json:"interfaceMode"`
			} `json:"vlanInformation"`
			BandWidth  int    `json:"bandwidth"`
			LinkStatus string `json:"linkStatus"`
		} `json:"Ethernet31/1"`
		Ethernet_32_1 struct {
			VlanInformation struct {
				VlanExplanation string `json:"vlanExplanation"`
				InterfaceMode   string `json:"interfaceMode"`
			} `json:"vlanInformation"`
			BandWidth  int    `json:"bandwidth"`
			LinkStatus string `json:"linkStatus"`
		} `json:"Ethernet32/1"`
		Ethernet_33_1 struct {
			VlanInformation struct {
				VlanExplanation string `json:"vlanExplanation"`
				InterfaceMode   string `json:"interfaceMode"`
			} `json:"vlanInformation"`
			BandWidth  int    `json:"bandwidth"`
			LinkStatus string `json:"linkStatus"`
		} `json:"Ethernet33/1"`
		Ethernet_34_1 struct {
			VlanInformation struct {
				VlanExplanation string `json:"vlanExplanation"`
				InterfaceMode   string `json:"interfaceMode"`
			} `json:"vlanInformation"`
			BandWidth  int    `json:"bandwidth"`
			LinkStatus string `json:"linkStatus"`
		} `json:"Ethernet34/1"`
		Ethernet_35_1 struct {
			VlanInformation struct {
				VlanExplanation string `json:"vlanExplanation"`
				InterfaceMode   string `json:"interfaceMode"`
			} `json:"vlanInformation"`
			BandWidth  int    `json:"bandwidth"`
			LinkStatus string `json:"linkStatus"`
		} `json:"Ethernet35/1"`
		Ethernet_36_1 struct {
			VlanInformation struct {
				VlanExplanation string `json:"vlanExplanation"`
				InterfaceMode   string `json:"interfaceMode"`
			} `json:"vlanInformation"`
			BandWidth  int    `json:"bandwidth"`
			LinkStatus string `json:"linkStatus"`
		} `json:"Ethernet36/1"`
		PortChannel1 struct {
			VlanInformation struct {
				VlanExplanation string `json:"vlanExplanation"`
				InterfaceMode   string `json:"interfaceMode"`
			} `json:"vlanInformation"`
			BandWidth  int    `json:"bandwidth"`
			LinkStatus string `json:"linkStatus"`
		} `json:"Port-Channel1"`
		PortChannel2 struct {
			VlanInformation struct {
				VlanExplanation string `json:"vlanExplanation"`
				InterfaceMode   string `json:"interfaceMode"`
			} `json:"vlanInformation"`
			BandWidth  int    `json:"bandwidth"`
			LinkStatus string `json:"linkStatus"`
		} `json:"Port-Channel2"`
		PortChannel3 struct {
			VlanInformation struct {
				VlanExplanation string `json:"vlanExplanation"`
				InterfaceMode   string `json:"interfaceMode"`
			} `json:"vlanInformation"`
			BandWidth  int    `json:"bandwidth"`
			LinkStatus string `json:"linkStatus"`
		} `json:"Port-Channel3"`
		PortChannel4 struct {
			VlanInformation struct {
				VlanExplanation string `json:"vlanExplanation"`
				InterfaceMode   string `json:"interfaceMode"`
			} `json:"vlanInformation"`
			BandWidth  int    `json:"bandwidth"`
			LinkStatus string `json:"linkStatus"`
		} `json:"Port-Channel4"`
		PortChannel5 struct {
			VlanInformation struct {
				VlanExplanation string `json:"vlanExplanation"`
				InterfaceMode   string `json:"interfaceMode"`
			} `json:"vlanInformation"`
			BandWidth  int    `json:"bandwidth"`
			LinkStatus string `json:"linkStatus"`
		} `json:"Port-Channel5"`
		PortChannel6 struct {
			VlanInformation struct {
				VlanExplanation string `json:"vlanExplanation"`
				InterfaceMode   string `json:"interfaceMode"`
			} `json:"vlanInformation"`
			BandWidth  int    `json:"bandwidth"`
			LinkStatus string `json:"linkStatus"`
		} `json:"Port-Channel6"`
		PortChannel7 struct {
			VlanInformation struct {
				VlanExplanation string `json:"vlanExplanation"`
				InterfaceMode   string `json:"interfaceMode"`
			} `json:"vlanInformation"`
			BandWidth  int    `json:"bandwidth"`
			LinkStatus string `json:"linkStatus"`
		} `json:"Port-Channel7"`
		PortChannel8 struct {
			VlanInformation struct {
				VlanExplanation string `json:"vlanExplanation"`
				InterfaceMode   string `json:"interfaceMode"`
			} `json:"vlanInformation"`
			BandWidth  int    `json:"bandwidth"`
			LinkStatus string `json:"linkStatus"`
		} `json:"Port-Channel8"`
		Management1 struct {
			VlanInformation struct {
				VlanExplanation string `json:"vlanExplanation"`
				InterfaceMode   string `json:"interfaceMode"`
			} `json:"vlanInformation"`
			BandWidth  int    `json:"bandwidth"`
			LinkStatus string `json:"linkStatus"`
		} `json:"Management1"`
	} `json:"interfaceStatuses"`
}
