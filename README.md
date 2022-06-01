# Yaml2go

Just another tool that convert YAML to GO struct.

## What's for:
Before:
```yaml
web: 172.16.20.30:19000
db:
  ch:
    url: 172.16.20.20:9000
    port: 19000
    ver: 19.09
    user: default
    passwd: somepassword
    some_list:
      - aaaaa
      - bbbb
```
After:
```go
package Config

type Config struct {
	Web string `yaml:"web"`
	Db  struct {
		Ch struct {
			Passwd   string   `yaml:"passwd"`
			Url      string   `yaml:"url"`
			Port     int      `yaml:"port"`
			Ver      float64  `yaml:"ver"`
			User     string   `yaml:"user"`
			SomeList []string `yaml:"some_list"`
		}
	}
}

```
See [example.go](example/example.go) for usage.
