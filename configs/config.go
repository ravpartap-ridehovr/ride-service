package config

import (
	fmt "fmt"
	log "log"
	os "os"
	strings "strings"

	config "github.com/ridehovr/goconfig"
)

// Configurations exported.
type Configurations struct {
	Service  ServiceConfigurations
	Redis    RedisConfigurations
	Broker   BrokerConfigurations
	Clients  []*ClientConfigurations
	Postgres Postgres
}

// ServerConfigurations exported.
type ServiceConfigurations struct {
	Name              string
	Port              string
	TwilioPhoneNumber string
	Secrets           ServiceSecretsConfiguratios
	AuthPublicKey     string
	AuthPrivateKey    string
}

// secrets for service.
type ServiceSecretsConfiguratios struct {
	TwilioUsername string
	TwilioAPIToken string

	AuthnPrivateKey string
}

type RedisConfigurations struct {
	Endpoint string
}

type BrokerConfigurations struct {
	Endpoint string
	Port     string
	Protocol string
	Secrets  BrokerSecretsConfigurations
}

type BrokerSecretsConfigurations struct {
	Username string
	Password string
}

type ClientConfigurations struct {
	Name      string
	Endpoint  string
	ProtoPath string
}

type Postgres struct {
	Dns string
}

// load configs.
func New() (*Configurations, error) {
	// Prefix for service and global vars.
	var (
		ServicePrefix = "REGISTRAR_"
		GlobalPrefix  = "GLOBAL_"
	)

	// vars to be loadded.
	var (
		Secrets = map[string]string{
			"SERVICE": ServicePrefix,
			"Broker":  GlobalPrefix,
		}
		Configs = map[string]string{
			"SERVICE": ServicePrefix,
			"BROKER":  GlobalPrefix,
			"AUTHN":   GlobalPrefix,
			"REDIS":   GlobalPrefix,
		}
	)

	env := os.Getenv("ENV")

	// new configurations.
	readConfigs := &Configurations{}
	// load from file.
	newConfig := config.New(readConfigs)

	// fix for wired race condition
	// if no env is set just read file
	if env == "" {
		// read file.
		err := newConfig.ReadYamlFile("config.yml")
		if err != nil {
			return nil, fmt.Errorf("unable to read config yml file: %w", err)
		}

		return readConfigs, nil
	}

	// read configs from parameter store
	for config, prefix := range Configs {
		// update parameter name.
		prefix = strings.ToUpper(env + "_" + prefix)
		config = strings.ToUpper(prefix + config)

		// read from parameter store.
		err := newConfig.ReadFromParameterStore(config, prefix)
		if err != nil {
			log.Printf("unable to read %s configs from parameter store: %v", config, err)
		}
	}

	// read global secrets and parameters.
	// read from secrets manager.
	for secret, prefix := range Secrets {
		// update var name.
		prefix = strings.ToUpper(env + "_" + prefix)
		secret = strings.ToUpper(prefix + secret + "_SECRETS")

		err := newConfig.ReadFromSecretManger(secret, prefix)
		if err != nil {
			log.Printf("unable to read %s secrets from secret manager: %v", secret, err)
		}
	}

	// read global vars.
	return readConfigs, nil
}
