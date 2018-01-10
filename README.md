# go-mirakurun

go-mirakurun is a [Mirakurun](https://github.com/Chinachu/Mirakurun) Client for Go.

## Usage

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ykzts/go-mirakurun/mirakurun"
)

func main() {
	client := mirakurun.NewClient()
	ctx := context.Background()

	channels, _, err := client.Channels.List(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	for _, channel := range channels {
		for _, service := range channel.Services {
			fmt.Printf("%d: %s\n", service.ID, service.Name)
		}
	}
}
```

## License

[MIT](LICENSE)
