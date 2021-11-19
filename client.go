package npbapi

//Client Initialize Constructor Variables
type Client struct {
	Hostname string
	Username string
	Password string
}

//NpbClient Initialize Constructor
func NpbClient(hostname string, username string, password string) *Client {

	return &Client{
		Hostname: hostname,
		Username: username,
		Password: password,
	}
}
