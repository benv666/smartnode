package config

// Defaults
const (
	defaultBitflyNodeMetricsSecret      string = ""
	defaultBitflyNodeMetricsEndpoint    string = "https://beaconcha.in/api/v1/client/metrics"
	defaultBitflyNodeMetricsMachineName string = "Smartnode"
)

// Configuration for Bitfly Node Metrics
type BitflyNodeMetricsConfig struct {
	Title string `yaml:"-"`

	Secret Parameter `yaml:"secret,omitempty"`

	Endpoint Parameter `yaml:"endpoint,omitempty"`

	MachineName Parameter `yaml:"machineName, omitempty"`
}

// Generates a new Bitfly Node Metrics config
func NewBitflyNodeMetricsConfig(config *RocketPoolConfig) *BitflyNodeMetricsConfig {
	return &BitflyNodeMetricsConfig{
		Title: "Bitfly Node Metrics Settings",

		Secret: Parameter{
			ID:                   "bitflySecret",
			Name:                 "Node Metrics Secret",
			Description:          "The secret used to authenticate your Beaconcha.in node metrics integration. Can be found in your Beaconcha.in account settings.\n\nPlease visit https://beaconcha.in/login to access your account information.",
			Type:                 ParameterType_String,
			Default:              map[Network]interface{}{Network_All: defaultBitflyNodeMetricsSecret},
			AffectsContainers:    []ContainerID{ContainerID_Validator, ContainerID_Eth2},
			EnvironmentVariables: []string{"BITFLY_NODE_METRICS_SECRET"},
			// ensures the string is 28 characters of Base64
			Regex:              "^[A-Za-z0-9+/]{28}$",
			CanBeBlank:         false,
			OverwriteOnUpgrade: false,
		},

		Endpoint: Parameter{
			ID:                   "bitflyEndpoint",
			Name:                 "Node Metrics Endpoint",
			Description:          "The endpoint to send your Beaconcha.in Node Metrics data to. Should be left as the default.",
			Type:                 ParameterType_String,
			Default:              map[Network]interface{}{Network_All: defaultBitflyNodeMetricsEndpoint},
			AffectsContainers:    []ContainerID{ContainerID_Validator, ContainerID_Eth2},
			EnvironmentVariables: []string{"BITFLY_NODE_METRICS_ENDPOINT"},
			CanBeBlank:           false,
			OverwriteOnUpgrade:   false,
		},

		MachineName: Parameter{
			ID:                   "bitflyMachineName",
			Name:                 "Node Metrics Machine Name",
			Description:          "The name of the machine you are running on. This is used to identify your machine in the mobile app.\nChange this if you are running multiple Smartnodes with the same Secret.",
			Type:                 ParameterType_String,
			Default:              map[Network]interface{}{Network_All: defaultBitflyNodeMetricsMachineName},
			AffectsContainers:    []ContainerID{ContainerID_Validator, ContainerID_Eth2},
			EnvironmentVariables: []string{"BITFLY_NODE_METRICS_MACHINE_NAME"},
			CanBeBlank:           false,
			OverwriteOnUpgrade:   false,
		},
	}
}

// Get the parameters for this config
func (config *BitflyNodeMetricsConfig) GetParameters() []*Parameter {
	return []*Parameter{
		&config.Secret,
		&config.Endpoint,
		&config.MachineName,
	}
}

// The the title for the config
func (config *BitflyNodeMetricsConfig) GetConfigTitle() string {
	return config.Title
}
