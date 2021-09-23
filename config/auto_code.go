package config

type Autocode struct {
	TransferRestart bool   `mapstructure:"transfer-restart" json:"transferRestart" yaml:"transfer-restart"`
	Root            string `mapstructure:"root" json:"root" yaml:"root"`
}
