package conf

type Postgres struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DB       string `mapstructure:"db"`
}

type Auth struct {
	Secret  string `mapstructure:"secret"`
	Expires int    `mapstructure:"expires"`
}

type Admin struct {
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}

type Conf struct {
	Postgres Postgres `mapstructure:"postgres"`
	Auth     Auth     `mapstructure:"auth"`
	Admin    Admin    `mapstructure:"admin"`
}
