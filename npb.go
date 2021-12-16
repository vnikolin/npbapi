package npbapi

import (
	"bytes"
	"encoding/json"
)

// Show NPB Version
func (client *Client) ShowVersion() (VersionInfo, error) {
	var version VersionInfo
	var readBytes bytes.Buffer

	session, err := client.ConnectSSH()
	if err != nil {
		return version, err
	}
	defer client.CloseSSH(session)
	session.Stdout = &readBytes
	err = session.Run("show version | json")
	if err != nil {
		return version, err
	}

	//json.Unmarshal([]byte(readBytes.String()), &version)
	json.Unmarshal(readBytes.Bytes(), &version)
	return version, nil
}

// Show NPB Host Info
func (client *Client) ShowHostname() (HostInfo, error) {
	var returnData HostInfo
	var readBytes bytes.Buffer

	session, err := client.ConnectSSH()
	if err != nil {
		return returnData, err
	}
	defer client.CloseSSH(session)
	session.Stdout = &readBytes
	err = session.Run("show hostname | json")
	if err != nil {
		return returnData, err
	}

	json.Unmarshal(readBytes.Bytes(), &returnData)
	return returnData, nil
}

// Show NPB Power Info
func (client *Client) ShowPower() (PowerInfo, error) {
	var rawData RawPowerInfo
	var returnData PowerInfo
	var readBytes bytes.Buffer

	session, err := client.ConnectSSH()
	if err != nil {
		return returnData, err
	}
	defer client.CloseSSH(session)
	session.Stdout = &readBytes
	err = session.Run("show environment power | json")
	if err != nil {
		return returnData, err
	}

	json.Unmarshal(readBytes.Bytes(), &rawData)

	returnData = PowerInfo{
		PS1: rawData.PowerSupplies.Instance1.State,
		PS2: rawData.PowerSupplies.Instance2.State,
	}

	return returnData, nil
}

// Show NPB NTP Info
func (client *Client) ShowNtp() (NtpInfo, error) {
	var returnData NtpInfo
	var readBytes bytes.Buffer

	session, err := client.ConnectSSH()
	if err != nil {
		return returnData, err
	}
	defer client.CloseSSH(session)
	session.Stdout = &readBytes
	err = session.Run("show ntp status | json")
	if err != nil {
		return returnData, err
	}

	json.Unmarshal(readBytes.Bytes(), &returnData)
	return returnData, nil
}

// Show Interface Info
func (client *Client) ShowInt() (IntInfo, error) {
	var returnData IntInfo
	var readBytes bytes.Buffer

	session, err := client.ConnectSSH()
	if err != nil {
		return returnData, err
	}
	defer client.CloseSSH(session)
	session.Stdout = &readBytes
	err = session.Run("show interfaces status | json")
	if err != nil {
		return returnData, err
	}

	json.Unmarshal(readBytes.Bytes(), &returnData)

	return returnData, nil
}
