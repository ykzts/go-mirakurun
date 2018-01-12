# go-mirakurun

go-mirakurun is a [Mirakurun](https://github.com/Chinachu/Mirakurun) Client for Go.

[![build status](https://travis-ci.org/ykzts/go-mirakurun.svg?branch=master)](https://travis-ci.org/ykzts/go-mirakurun) [![GoDoc](https://godoc.org/github.com/ykzts/go-mirakurun/mirakurun?status.svg)](https://godoc.org/github.com/ykzts/go-mirakurun/mirakurun)

## Usage

```go
import "github.com/ykzts/go-mirakurun/mirakurun"
```

### Recoding

```go
ctx := context.Background()
ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
defer cancel()

name := fmt.Sprintf("stream-%d.ts", time.Now().Unix())
file, err := os.Create(name)
if err != nil {
        log.Fatal(err)
}
defer file.Close()

c := mirakurun.NewClient()

stream, _, err := c.GetServiceStream(ctx, 3239123608, true)
if err != nil {
        log.Fatal(err)
}
defer stream.Close()

io.Copy(file, stream)
```

[GoDoc](https://godoc.org/github.com/ykzts/go-mirakurun/mirakurun)

## License

[MIT](LICENSE)
