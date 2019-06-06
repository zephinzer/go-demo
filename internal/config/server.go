package config

/*

For use in READMEs:

| Environment Variable | Description |
| --- | --- |
| `HOST` | Sets the host interface to bind to (defaults to 0.0.0.0) |
| `PORT` | Sets the port to listen on (defaults to 11111) |

*/

import (
	"fmt"
	"github.com/spf13/viper"
)

func NewServer() Server {
	config := viper.New()
	config.SetDefault("host", "0.0.0.0")
	config.SetDefault("port", 11111)
	config.AutomaticEnv()
	return Server{
		Host: config.GetString("host"),
		Port: uint16(config.GetInt("port")),
	}
}

type Server struct {
	Host string
	Port uint16
}

// GetAddr returns a usable Addr for the http.Server struct
func (server Server) GetAddr() string {
	return fmt.Sprintf("%s:%v", server.Host, server.Port)
}