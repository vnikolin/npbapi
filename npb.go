package npbapi

import (
	"bytes"
	"encoding/json"
	"strings"

	"golang.org/x/crypto/ssh"
)

// Open SSH Connection
func (c *Client) ConnectSSH() (*ssh.Session, error) {
	hostname := c.Hostname + ":22"
	config := &ssh.ClientConfig{
		User: c.Username,
		Auth: []ssh.AuthMethod{
			ssh.KeyboardInteractive(func(user, instruction string, questions []string, echos []bool) ([]string, error) {
				if len(questions) == 1 && strings.Contains(questions[0], "Password") {
					return []string{string(c.Password)}, nil
				}
				return []string{}, nil
			}),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	client, err := ssh.Dial("tcp", hostname, config)
	if err != nil {
		return nil, err
	}
	session, err := client.NewSession()
	if err != nil {
		return nil, err
	}
	return session, nil
}

// Close SSH Connection
func (c *Client) CloseSSH(session *ssh.Session) {
	session.Close()
}

// Show NPB Version
func (c *Client) ShowVersion(session *ssh.Session) (VersionInfo, error) {
	var version VersionInfo
	var readBytes bytes.Buffer

	session.Stdout = &readBytes
	err := session.Run("show version | json")
	if err != nil {
		return version, err
	}

	//json.Unmarshal([]byte(readBytes.String()), &version)
	json.Unmarshal(readBytes.Bytes(), &version)
	return version, nil
}

// Show NPB Host Info
func (c *Client) ShowHostname(session *ssh.Session) (HostInfo, error) {
	var returnData HostInfo
	var readBytes bytes.Buffer

	session.Stdout = &readBytes
	err := session.Run("show hostname | json")
	if err != nil {
		return returnData, err
	}

	json.Unmarshal(readBytes.Bytes(), &returnData)
	return returnData, nil
}

// Show NPB Power Info
func (c *Client) ShowPower(session *ssh.Session) (PowerInfo, error) {
	var rawData RawPowerInfo
	var returnData PowerInfo
	var readBytes bytes.Buffer

	session.Stdout = &readBytes
	err := session.Run("show environment power | json")
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
func (c *Client) ShowNtp(session *ssh.Session) (NtpInfo, error) {
	var returnData NtpInfo
	var readBytes bytes.Buffer

	session.Stdout = &readBytes
	err := session.Run("show ntp status | json")
	if err != nil {
		return returnData, err
	}

	json.Unmarshal(readBytes.Bytes(), &returnData)
	return returnData, nil
}

// Show Interface Info
func (c *Client) ShowInt(session *ssh.Session) (IntInfo, error) { //(IntInfo, error) {
	//var rawData RawIntInfo
	var returnData IntInfo
	var readBytes bytes.Buffer

	session.Stdout = &readBytes
	err := session.Run("show interfaces status | json")
	if err != nil {
		return returnData, err
	}

	//fmt.Println(readBytes.String())
	json.Unmarshal(readBytes.Bytes(), &returnData)
	//fmt.Printf("%+v\n", returnData)
	//fmt.Printf("%#v\n", returnData)

	return returnData, nil
}
