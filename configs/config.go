package configs

import (
	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

type conf struct {
	DBDriver   string `mapstructure:"DB_DRIVER"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`

	WebServerPort string `mapstructure:"WEB_SERVER_PORT"`

	JWTSecret   string `mapstructure:"JWT_SECRET"`
	JWTExpireIn int    `mapstructure:"JWT_EXPIRE_IN"`

	TokenAuth *jwtauth.JWTAuth
}

func LoadConfig(path string) (*conf, error) {
	var cfg *conf

	viper.SetConfigName("app_config") // name of config file (without extension)
	viper.SetConfigType("env")        // type of config file
	viper.AddConfigPath(path)         // path to look for the config file in
	viper.SetConfigFile(".env")       // config file (without extension) (name of config file)
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}

	cfg.TokenAuth = jwtauth.New("HS256", []byte(cfg.JWTSecret), nil)

	return cfg, nil
}
