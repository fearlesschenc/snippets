package viper

import "github.com/spf13/viper"

type Foo struct {
	A int `yaml:"a,omitempty"`
}

type Config struct {
	//Foo *Foo `yaml:"foo,omitempty"`
	Foosdfasff *Foo `json:"foo" yaml:"foo"`
}

func Unmarshal(val interface{}) error {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	return viper.Unmarshal(val)
}

func Main() {
	config := &Config{
		Foosdfasff: &Foo{A: 11},
	}
	Unmarshal(config)

	print(config.Foosdfasff.A)
}
