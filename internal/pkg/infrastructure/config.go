package config

type Config struct {
	ServerConfig
	JsonSchemaConfig

	ApplicationName string `env:"APPLICATION_NAME"`
	AppEnv          string `env:"APP_ENV"`
}

type ServerConfig struct {
	ServerHost         string `env:"SERVER_HOST"`
	ServerPort         string `env:"SERVER_PORT"`
	ServerWriteTimeout int    `env:"SERVER_WRITE_TIMEOUT"`
	ServerReadTimeout  int    `env:"SERVER_READ_TIMEOUT"`
}

type JsonSchemaConfig struct {
	SchemaPath string `env:"BASE_JSON_SCHEMA_PATH"`
}
