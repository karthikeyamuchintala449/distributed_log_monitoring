package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type Config struct {
	// App cnfigurations
	AppName    string
	AppVersion string
	Debug      bool
	Environment string

	// API configurations

	APIHost    string
	APIPort    int
	APIPrefix  string

	// Elasticsearch configuratiosn for the elsatic search
	ElasticsearchHost       string
	ElasticsearchPort       int
	ElasticsearchScheme     string
	ElasticsearchIndexPrefix string
	ElasticsearchShards     int
	ElasticsearchReplicas   int
	ElasticsearchUsername   string
	ElasticsearchPassword   string

	// Kafka service configuraions 
	KafkaBrokers       []string
	KafkaTopicLogs     string
	KafkaTopicAlerts   string
	KafkaConsumerGroup string
	KafkaBatchSize     int

	// Logging configurations

	LogLevel string
	LogFormat string // json or text

	// Alerting service configurations for alerting  and notification
	AlertEnabled      bool
	EmailEnabled      bool
	EmailSMTPHost     string
	EmailSMTPPort     int
	EmailSender       string
	SlackWebhookURL   string

	// Features toggles for enabling or disabling features
	AnomalyDetectionEnabled  bool
	RealTimeStreamingEnabled bool

	// Performance configuraions
	MaxQueryResults     int
	SearchTimeoutSeconds int
	BatchProcessingSize int
	WorkerPoolSize      int
	GracefulShutdownTimeout time.Duration
}

// LoadConfig loads configuration from environment variables
func LoadConfig() *Config {
	cfg := &Config{
		// App defaults
		AppName:    getEnv("APP_NAME", "Distributed Log Monitoring System"),
		AppVersion: getEnv("APP_VERSION", "2.0.0"),
		Debug:      getEnvBool("DEBUG", false),
		Environment: getEnv("ENVIRONMENT", "development"),

		// API defaults
		APIHost:   getEnv("API_HOST", "0.0.0.0"),
		APIPort:   getEnvInt("API_PORT", 8000),
		APIPrefix: getEnv("API_PREFIX", "/api/v1"),

		// Elasticsearch defaults
		ElasticsearchHost:        getEnv("ELASTICSEARCH_HOST", "localhost"),
		ElasticsearchPort:        getEnvInt("ELASTICSEARCH_PORT", 9200),
		ElasticsearchScheme:      getEnv("ELASTICSEARCH_SCHEME", "http"),
		ElasticsearchIndexPrefix: getEnv("ELASTICSEARCH_INDEX_PREFIX", "logs"),
		ElasticsearchShards:      getEnvInt("ELASTICSEARCH_SHARDS", 3),
		ElasticsearchReplicas:    getEnvInt("ELASTICSEARCH_REPLICAS", 1),
		ElasticsearchUsername:    getEnv("ELASTICSEARCH_USERNAME", ""),
		ElasticsearchPassword:    getEnv("ELASTICSEARCH_PASSWORD", ""),

		// Kafka defaults
		KafkaTopicLogs:     getEnv("KAFKA_TOPIC_LOGS", "logs"),
		KafkaTopicAlerts:   getEnv("KAFKA_TOPIC_ALERTS", "alerts"),
		KafkaConsumerGroup: getEnv("KAFKA_CONSUMER_GROUP", "log_monitor"),
		KafkaBatchSize:     getEnvInt("KAFKA_BATCH_SIZE", 1000),

		// Logging defaults
		LogLevel:  getEnv("LOG_LEVEL", "info"),
		LogFormat: getEnv("LOG_FORMAT", "json"),

		// Alerting defaults
		AlertEnabled:    getEnvBool("ALERT_ENABLED", true),
		EmailEnabled:    getEnvBool("EMAIL_ENABLED", false),
		EmailSMTPHost:   getEnv("EMAIL_SMTP_HOST", "smtp.gmail.com"),
		EmailSMTPPort:   getEnvInt("EMAIL_SMTP_PORT", 587),
		EmailSender:     getEnv("EMAIL_SENDER", "noreply@monitoring.local"),
		SlackWebhookURL: getEnv("SLACK_WEBHOOK_URL", ""),

		// Features defaults
		AnomalyDetectionEnabled:  getEnvBool("ANOMALY_DETECTION_ENABLED", true),
		RealTimeStreamingEnabled: getEnvBool("REAL_TIME_STREAMING_ENABLED", true),

		// Performance defaults
		MaxQueryResults:         getEnvInt("MAX_QUERY_RESULTS", 10000),
		SearchTimeoutSeconds:    getEnvInt("SEARCH_TIMEOUT_SECONDS", 30),
		BatchProcessingSize:     getEnvInt("BATCH_PROCESSING_SIZE", 5000),
		WorkerPoolSize:          getEnvInt("WORKER_POOL_SIZE", 10),
		GracefulShutdownTimeout: time.Duration(getEnvInt("GRACEFUL_SHUTDOWN_TIMEOUT_SECONDS", 30)) * time.Second,
	}

	// Parse Kafka brokers
	kafkaBrokers := getEnv("KAFKA_BROKERS", "localhost:9092")
	cfg.KafkaBrokers = parseStringList(kafkaBrokers)

	return cfg
}

// Helper functions
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
	if value, exists := os.LookupEnv(key); exists {
		if intVal, err := strconv.Atoi(value); err == nil {
			return intVal
		}
	}
	return defaultValue
}

func getEnvBool(key string, defaultValue bool) bool {
	if value, exists := os.LookupEnv(key); exists {
		return value == "true" || value == "1" || value == "yes"
	}
	return defaultValue
}

func parseStringList(input string) []string {
	if input == "" {
		return []string{}
	}
	// Simple comma-separated parsing
	var result []string
	for _, item := range splitString(input, ",") {
		if trimmed := trimSpace(item); trimmed != "" {
			result = append(result, trimmed)
		}
	}
	return result
}

func splitString(s, sep string) []string {
	var result []string
	var current string
	for i := 0; i < len(s); i++ {
		if string(s[i]) == sep {
			result = append(result, current)
			current = ""
		} else {
			current += string(s[i])
		}
	}
	if current != "" {
		result = append(result, current)
	}
	return result
}

func trimSpace(s string) string {
	for len(s) > 0 && (s[0] == ' ' || s[0] == '\t' || s[0] == '\n' || s[0] == '\r') {
		s = s[1:]
	}
	for len(s) > 0 && (s[len(s)-1] == ' ' || s[len(s)-1] == '\t' || s[len(s)-1] == '\n' || s[len(s)-1] == '\r') {
		s = s[:len(s)-1]
	}
	return s
}

// GetElasticsearchURL returns the Elasticsearch URL
func (c *Config) GetElasticsearchURL() string {
	return fmt.Sprintf("%s://%s:%d", c.ElasticsearchScheme, c.ElasticsearchHost, c.ElasticsearchPort)
}

// IsProduction returns true if environment is production
func (c *Config) IsProduction() bool {
	return c.Environment == "production"
}
