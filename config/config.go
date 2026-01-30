package config

type Server struct {
	Zap    Zap    `mapstructure:"zap" json:"zap" yaml:"zap"`
	System System `mapstructure:"system" json:"system" yaml:"system"`

	// auto
	AutoCode Autocode        `mapstructure:"autocode" json:"autocode" yaml:"autocode"`
	DBList   []SpecializedDB `mapstructure:"db-list" json:"db-list" yaml:"db-list"`
}
