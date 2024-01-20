package config

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	log "github.com/cjlapao/common-go-logger"
	"github.com/cjlapao/common-go/helper"
	"github.com/cjlapao/github-templater/pkg/constants"
	"github.com/cjlapao/github-templater/pkg/context"
	"gopkg.in/yaml.v3"
)

var globalConfig *Config

var extensions = []string{
	".local.yaml",
	".local.yml",
	".local.json",
	".yaml",
	".yml",
	".json",
}

type Config struct {
	ctx    *context.ProvisionerContext
	mode   string
	config map[string]string
}

func New(ctx *context.ProvisionerContext) *Config {
	globalConfig = &Config{
		mode:   "api",
		ctx:    ctx,
		config: make(map[string]string),
	}

	globalConfig.LogLevel()
	return globalConfig
}

func Get() *Config {
	if globalConfig == nil {
		return New(context.New())
	}

	return globalConfig
}

func (c *Config) Load() bool {
	fileName := ""
	configFileName := helper.GetFlagValue(constants.ConfigFileFlag, "")
	if configFileName != "" {
		if _, err := os.Stat(configFileName); !os.IsNotExist(err) {
			fileName = configFileName
		}
	} else {
		for _, extension := range extensions {
			if _, err := os.Stat(fmt.Sprintf("config%s", extension)); !os.IsNotExist(err) {
				fileName = fmt.Sprintf("config%s", extension)
				break
			}
		}
	}

	if fileName == "" {
		c.ctx.LogInfo("No configuration file found")
		return false
	}

	c.ctx.LogInfo("Loading configuration from %s", fileName)
	content, err := helper.ReadFromFile(fileName)
	if err != nil {
		c.ctx.LogError("Error reading configuration file: %s", err.Error())
		return false
	}

	if content == nil {
		c.ctx.LogError("Error reading configuration file: %s", err.Error())
		return false
	}

	if strings.HasSuffix(fileName, ".json") {
		err = json.Unmarshal(content, &c.config)
		if err != nil {
			c.ctx.LogError("Error reading configuration file: %s", err.Error())
			return false
		}
	} else {
		err = yaml.Unmarshal(content, &c.config)
		if err != nil {
			c.ctx.LogError("Error reading configuration file: %s", err.Error())
			return false
		}
	}

	return true
}

func (c *Config) LogLevel() string {
	logLevel := c.GetKey(constants.EnvLogLevel)
	if logLevel != "" {
		c.ctx.LogInfo("Log Level set to %v", logLevel)
	}
	switch strings.ToLower(logLevel) {
	case "debug":
		logLevel = "DEBUG"
		c.ctx.Logger().LogLevel = log.Debug
	case "info":
		logLevel = "INFO"
		c.ctx.Logger().LogLevel = log.Info
	case "warn":
		logLevel = "WARN"
		c.ctx.Logger().LogLevel = log.Warning
	case "error":
		logLevel = "ERROR"
		c.ctx.Logger().LogLevel = log.Error
	}

	return logLevel
}

func (c *Config) GetKey(key string) string {
	value := helper.GetFlagValue(key, "")
	exists := false

	if value == "" {
		value, exists = os.LookupEnv(key)
		if value == "" && !exists {
			for k, v := range c.config {
				if strings.EqualFold(k, key) {
					value = v
					break
				}
			}
		}
	}

	return strings.TrimSpace(value)
}

func (c *Config) GetIntKey(key string) int {
	value := c.GetKey(key)
	if value == "" {
		return 0
	}

	intVal, err := strconv.Atoi(value)
	if err != nil {
		return 0
	}

	return intVal
}

func (c *Config) GetBoolKey(key string) bool {
	value := c.GetKey(key)
	if value == "" {
		return false
	}

	boolVal, err := strconv.ParseBool(value)
	if err != nil {
		return false
	}

	return boolVal
}

func (c *Config) RequestFromUser(key string, question string, defaultValue string) string {
	value := c.GetKey(key)
	if value != "" {
		return value
	}

	fmt.Printf("%s: ", question)
	fmt.Scanln(&value)
	c.ctx.LogDebug("User input: %v", value)
	if value == "" {
		return defaultValue
	}

	return value
}

func (c *Config) RequestBoolFromUser(key string, question string, defaultValue bool) bool {
	value := c.GetKey(key)
	if value == "" {
		fmt.Printf("%s: ", question)
		fmt.Scanln(&value)
		if value == "" {
			c.ctx.LogDebug("User input: %v, could not convert it to bool", value)
			return defaultValue
		}
	}

	if strings.EqualFold(value, "y") || strings.EqualFold(value, "yes") {
		value = "true"
	} else if strings.EqualFold(value, "n") || strings.EqualFold(value, "no") {
		value = "false"
	}

	boolVal, err := strconv.ParseBool(value)
	if err != nil {
		c.ctx.LogDebug("User input: %v, could not convert it to bool", value)
		return defaultValue
	}

	c.ctx.LogDebug("User input: %v, we converted to %v", value, boolVal)
	return boolVal
}

func (c *Config) RequestIntFromUser(key string, question string, defaultValue int) int {
	value := c.GetKey(key)
	if value == "" {
		fmt.Printf("%s: ", question)
		fmt.Scanln(&value)
		if value == "" {
			c.ctx.LogDebug("User input: %v, could not convert it to int", value)
			return defaultValue
		}
	}

	intVal, err := strconv.Atoi(value)
	if err != nil {
		c.ctx.LogDebug("User input: %v, could not convert it to int", value)
		return defaultValue
	}

	c.ctx.LogDebug("User input: %v, we converted to %v", value, intVal)
	return intVal
}
