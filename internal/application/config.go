package application

type screenName string

type ConfigApp struct {
	AppHeight float32 `yaml:"app-height"`
	AppWidth  float32 `yaml:"app-width"`
	Name      string  `yaml:"name"`
	FixedSize bool    `yaml:"fixed-size"`
}

// managerService map[string]interface{}

func NewConfig(appHeight float32, appWidth float32, name string, fixedsize bool) *ConfigApp {
	config := &ConfigApp{
		AppHeight: appHeight,
		AppWidth:  appWidth,
		Name:      name,
		FixedSize: fixedsize,
	}

	return config
}
