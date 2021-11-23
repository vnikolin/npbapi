package npbapi

import (
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
