# Arista Network Packet Broker API

This Library tested with Arista DCS-7280CR3K-32D4-R

## Usage

```go
//main.go

package main

import (
	"bytes"
	"fmt"

	"github.com/vnikolin/npbapi"
)

func main() {

	client := npbapi.NpbClient("x.x.x.x", "userid", "password")

	version, err := npbapi.ShowVersion(client)
	fmt.Printf("%#v\n", version)
	fmt.Println(err)

}