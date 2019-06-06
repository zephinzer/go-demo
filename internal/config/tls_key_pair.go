package config

/*

For use in READMEs:

| Environment Variable | Description |
| --- | --- |
| `CERT` | Absolute or relative path to the server certificate file (defaults to `/etc/ssl/server.crt`) |
| `KEY` | Absolute or relative path to the server key file (defaults to `/etc/ssl/server.key`) |

*/

import (
	"os"
	"path"

	"github.com/spf13/viper"
)

func NewTLSKeyPair() TLSKeyPair {
	config := viper.New()
	config.SetDefault("cert", "/etc/ssl/server.crt")
	config.SetDefault("key", "/etc/ssl/server.key")
	config.AutomaticEnv()

	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	certPath := config.GetString("cert")
	if !path.IsAbs(certPath) {
		certPath = path.Join(cwd, certPath)
	}

	keyPath := config.GetString("key")
	if !path.IsAbs(keyPath) {
		keyPath = path.Join(cwd, keyPath)
	}
	
	return TLSKeyPair{
		CertPath: certPath,
		KeyPath: keyPath,
	}
}

type TLSKeyPair struct {
	CertPath string
	KeyPath string
}