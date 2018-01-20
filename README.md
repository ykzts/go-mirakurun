# go-mirakurun

go-mirakurun is a [Mirakurun](https://github.com/Chinachu/Mirakurun) Client for Go.

[![build status](https://travis-ci.org/ykzts/go-mirakurun.svg?branch=master)](https://travis-ci.org/ykzts/go-mirakurun) [![GoDoc](https://godoc.org/github.com/ykzts/go-mirakurun/mirakurun?status.svg)](https://godoc.org/github.com/ykzts/go-mirakurun/mirakurun)

## Usage

```go
import "github.com/ykzts/go-mirakurun/mirakurun"
```

### Channel Scan

```go
c := mirakurun.NewClient()

opt := &mirakurun.ChannelScanOptions{Type: "BS"}
stream, _, err := c.ChannelScan(context.Background(), opt)
if err != nil {
        log.Fatal(err)
}
defer stream.Close()

io.Copy(os.Stdout, stream)
```

### Get Channel List

```go
c := mirakurun.NewClient()

channels, _, err := c.GetChannels(context.Background(), nil)
if err != nil {
        log.Fatal(err)
}

for _, channel := range channels {
        fmt.Printf("%s (%s): %s\n", channel.Channel, channel.Type, channel.Name)
}
```

### Recoding

```go
c := mirakurun.NewClient()

ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
defer cancel()

stream, _, err := c.GetServiceStream(ctx, 3239123608, true)
if err != nil {
        log.Fatal(err)
}
defer stream.Close()

name := fmt.Sprintf("stream-%d.ts", time.Now().Unix())
file, err := os.Create(name)
if err != nil {
        log.Fatal(err)
}
defer file.Close()

fmt.Printf("save to %s\n", name)

io.Copy(file, stream)
```

## License

[MIT](LICENSE)
