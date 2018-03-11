package cliutils

import (
	"io"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

// Config stores general setting params and provides accessors for them.
type Config struct {
	cfgFile   string
	v         *viper.Viper
	InReader  io.Reader
	OutWriter io.Writer
	ErrWriter io.Writer
	data      map[string]interface{}
}

// NewConfig creates new Config object.
func NewConfig(
	inReader io.Reader,
	outWriter io.Writer,
	errWriter io.Writer,
) *Config {
	return &Config{
		v:         viper.New(),
		InReader:  inReader,
		OutWriter: outWriter,
		ErrWriter: errWriter,
	}
}

// Init initializes the config read from the file.
func (c *Config) Init(cfgFile string) (err error) {
	c.cfgFile = cfgFile
	c.v.SetConfigFile(c.cfgFile)
	err = errors.WithStack(c.v.ReadInConfig())
	if err == nil {
		c.data = make(map[string]interface{})
		err = errors.WithStack(c.v.Unmarshal(&(c.data)))
	}
	return
}

// Name returns the application name.
func (c *Config) Name() (string, error) {
	v, err := c.templateValue("name")
	return v, errors.WithStack(err)
}

// Version returns the application version.
func (c *Config) Version() (string, error) {
	v, err := c.templateValue("version")
	return v, errors.WithStack(err)
}

func (c *Config) templateValue(key string) (string, error) {
	v, err := TemplateString(c.v.GetString(key)).Compile(c.data)
	return v, errors.WithStack(err)
}
