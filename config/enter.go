package config

type Config struct {
	Mysql  Mysql  `yaml:"mysql" json:"mysql"`
	System System `yaml:"system" json:"system"`
	Redis  Redis  `yaml:"initialize" json:"initialize"`
}
