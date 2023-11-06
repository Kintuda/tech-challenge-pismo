package config

type ServerConfig struct {
	Env             string `env:"ENV"`
	HttpPort        string `env:"HTTP_PORT,required"`
	PostgresDns     string `env:"POSTGRES_DNS,required"`
	EnableTracing   string `env:"ENABLE_TRACING"`
	CollectorUrl    string `env:"OTEL_EXPORTER_OTLP_ENDPOINT"`
	TracingInsecure string `env:"INSECURE_MODE"`
}
