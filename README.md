# go-mirakurun

## Usage

```go
import "github.com/ykzts/go-mirakurun/mirakurun"
```

```go
client := mirakurun.NewClient()
client.Host = "localhost"
client.Port = 40772

version, err := client.CheckVersion()
fmt.Println(version.Current)
```

## License

[MIT](LICENSE)
