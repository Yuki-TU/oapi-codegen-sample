package config

import (
	"github.com/caarlos0/env/v6"
)

type Config struct {
	Env               string `env:"GO_ENV" envDefault:"production"`
	Port              int    `env:"PORT" envDefault:"80"`
	DBHost            string `env:"DB_HOST" envDefault:"db"`
	DBPort            int    `env:"DB_PORT" envDefault:"3306"`
	DBUser            string `env:"DB_USER" envDefault:"admin"`
	DBPassword        string `env:"DB_PASSWORD" envDefault:"password"`
	DBName            string `env:"DB_NAME" envDefault:"point_app"`
	RedisHost         string `env:"REDIS_HOST" envDefault:"cache"`
	RedisPort         int    `env:"REDIS_PORT" envDefault:"6379"`
	AWSEndpoint       string `env:"AWS_ENDPOINT" envDefault:""`
	AWSId             string `env:"AWS_ACCESS_KEY_ID" envDefault:"accesskey"`
	AWSSecret         string `env:"AWS_SECRET_KEY" envDefault:"secretkey"`
	AWSRegion         string `env:"AWS_REGION" envDefault:"ap-northeast-1"`
	SenderMailAddress string `env:"SENDER_MAIL_ADDRESS" envDefault:"sample@sample.com"`
	FrontEndpoint     string `env:"FRONT_ENDPOINT" envDefault:"http://localhost:3000"`
}

var config Config

// Load は、環境変数から設定を読み込みます。
// 1回のみ呼び出してください。
func Load() error {
	if err := env.Parse(&config); err != nil {
		return err
	}
	return nil
}

func Get() Config { return config }

// func Load(ctx context.Context) (Config, error) {
// 	envconfig.MustProcess("", &config)

// 	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
// 	defer cancel()
// 	return config, nil
// }
