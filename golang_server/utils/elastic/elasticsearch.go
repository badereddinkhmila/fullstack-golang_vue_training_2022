package elastic

import (
	"com.fullstack.ecommerce/utils/config"
	"github.com/elastic/go-elasticsearch/v7"
	"log"
	"sync"
)

var (
	elasticSyncOnce sync.Once
	esClient        *elasticsearch.Client
)

func connect() {
	cfg := config.ElasticConfig()
	configuration := elasticsearch.Config{
		Addresses: cfg.Addresses,
		Username:  cfg.Username,
		Password:  cfg.Password,
	}

	es, err := elasticsearch.NewClient(configuration)

	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	res, err := es.Info()
	if err != nil {
		log.Fatalf("Error connecting with elasticsearch: %s", err)
	}
	defer res.Body.Close()
	esClient = es
}

func Client() *elasticsearch.Client {
	elasticSyncOnce.Do(func() {
		connect()
	})
	return esClient
}
