# eXTReMe IP Lookup Go Client

A simple API Client written in Go for https://extreme-ip-lookup.com/

## Examples

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/electrologue/extremeiplookup"
)

func main() {
	client := extremeiplookup.NewClient("secret")

	ipInfo, err := client.Lookup(context.Background(), "10.10.10.10")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(ipInfo)
}
```
