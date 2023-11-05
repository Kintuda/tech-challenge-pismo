package config

type ServerConfig struct {
	Env             string `env:"ENV"`
	LocalBucket     string `env:"LOCAL_BUCKET"`
	KYCBucket       string `env:"KYC_BUCKET"`
	HttpPort        string `env:"HTTP_PORT,required"`
	PostgresDns     string `env:"POSTGRES_DNS,required"`
	RedisUrl        string `env:"REDIS_URL"`
	RedisPassword   string `env:"REDIS_PASSWORD"`
	RedisDB         int    `env:"REDIS_DB"`
	EnableTracing   string `env:"ENABLE_TRACING"`
	CollectorUrl    string `env:"OTEL_EXPORTER_OTLP_ENDPOINT"`
	TracingInsecure string `env:"INSECURE_MODE"`
}
