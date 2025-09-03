package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoConfig struct {
	Host            string
	Port            int
	Username        string
	Password        string
	Database        string
	AuthSource      string
	ConnectTimeout  time.Duration
	MaxPoolSize     uint64
	MinPoolSize     uint64
	MaxConnIdleTime time.Duration
	ReplicaSet      string
	SSL             bool
}

type MongoDB struct {
	Client   *mongo.Client
	Database *mongo.Database
	Config   *MongoConfig
}

func LoadMongoConfig() *MongoConfig {
	config := &MongoConfig{
		Host:            getEnvMongoDb("MONGO_HOST", "localhost"),
		Port:            getEnvAsInt("MONGO_PORT", 27017),
		Username:        getEnvMongoDb("MONGO_USERNAME", ""),
		Password:        getEnvMongoDb("MONGO_PASSWORD", ""),
		Database:        getEnvMongoDb("MONGO_DATABASE", "myapp"),
		AuthSource:      getEnvMongoDb("MONGO_AUTH_SOURCE", "admin"),
		ConnectTimeout:  getEnvAsDuration("MONGO_CONNECT_TIMEOUT", 10*time.Second),
		MaxPoolSize:     getEnvAsUint64("MONGO_MAX_POOL_SIZE", 100),
		MinPoolSize:     getEnvAsUint64("MONGO_MIN_POOL_SIZE", 5),
		MaxConnIdleTime: getEnvAsDuration("MONGO_MAX_CONN_IDLE_TIME", 30*time.Second),
		ReplicaSet:      getEnvMongoDb("MONGO_REPLICA_SET", ""),
		SSL:             getEnvAsBool("MONGO_SSL", false),
	}

	return config
}

func ConnectMongoDB(config *MongoConfig) (*MongoDB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), config.ConnectTimeout)
	defer cancel()

	uri := buildMongoURI(config)

	clientOptions := options.Client().
		ApplyURI(uri).
		SetMaxPoolSize(config.MaxPoolSize).
		SetMinPoolSize(config.MinPoolSize).
		SetMaxConnIdleTime(config.MaxConnIdleTime)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB: %w", err)
	}

	database := client.Database(config.Database)

	log.Printf("Successfully connected to MongoDB database: %s", config.Database)

	return &MongoDB{
		Client:   client,
		Database: database,
		Config:   config,
	}, nil
}

func (m *MongoDB) Disconnect(ctx context.Context) error {
	if err := m.Client.Disconnect(ctx); err != nil {
		return fmt.Errorf("failed to disconnect from MongoDB: %w", err)
	}
	log.Println("Disconnected from MongoDB")
	return nil
}

func (m *MongoDB) GetCollection(name string) *mongo.Collection {
	return m.Database.Collection(name)
}

func (m *MongoDB) HealthCheck(ctx context.Context) error {
	if err := m.Client.Ping(ctx, readpref.Primary()); err != nil {
		return fmt.Errorf("MongoDB health check failed: %w", err)
	}
	return nil
}

func buildMongoURI(config *MongoConfig) string {
	var uri string

	if config.Username != "" && config.Password != "" {
		uri = fmt.Sprintf("mongodb://%s:%s@%s:%d",
			config.Username, config.Password, config.Host, config.Port)
	} else {
		uri = fmt.Sprintf("mongodb://%s:%d", config.Host, config.Port)
	}

	params := []string{}

	if config.AuthSource != "" && config.Username != "" {
		params = append(params, fmt.Sprintf("authSource=%s", config.AuthSource))
	}

	if config.ReplicaSet != "" {
		params = append(params, fmt.Sprintf("replicaSet=%s", config.ReplicaSet))
	}

	if config.SSL {
		params = append(params, "ssl=true")
	}

	if len(params) > 0 {
		uri += "?" + joinParams(params)
	}

	return uri
}

func getEnvMongoDb(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

func getEnvAsUint64(key string, defaultValue uint64) uint64 {
	if value := os.Getenv(key); value != "" {
		if uint64Value, err := strconv.ParseUint(value, 10, 64); err == nil {
			return uint64Value
		}
	}
	return defaultValue
}

func getEnvAsDuration(key string, defaultValue time.Duration) time.Duration {
	if value := os.Getenv(key); value != "" {
		if duration, err := time.ParseDuration(value); err == nil {
			return duration
		}
	}
	return defaultValue
}

func getEnvAsBool(key string, defaultValue bool) bool {
	if value := os.Getenv(key); value != "" {
		if boolValue, err := strconv.ParseBool(value); err == nil {
			return boolValue
		}
	}
	return defaultValue
}

func joinParams(params []string) string {
	result := ""
	for i, param := range params {
		if i > 0 {
			result += "&"
		}
		result += param
	}
	return result
}
