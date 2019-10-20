package config

// ES holds the configuration values for the Elasticsearch client.
type ES struct {
	Host    string `env:"ES_HOST" default:"192.168.178.26"`
	Port    string `env:"ES_PORT" default:"9200"`
	Cluster string `env:"ES_CLUSTER" default:"superheromatch"`
	Index   string `env:"ES_INDEX" default:"superhero"`
}