package model

// Reusable config goes here
type AppConfig struct {
	Encryption Encryption
}

type Encryption struct {
	Key string `mapstructure:"key"`
}
