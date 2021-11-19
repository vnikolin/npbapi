package npbapi

//Client Initialize the Constructor Variables
type Client struct {
	Hostname string
	Username string
	Password string
}

//NetworkClient Initialize the Constructor
func NpbClient(hostname string, username string, password string) *Client {

	return &Client{
		Hostname: hostname,
		Username: username,
		Password: password,
	}
}
