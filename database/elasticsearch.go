package database

import (
	"database/sql"
	"log"

	"github.com/elastic/go-elasticsearch/v8"
)

type EmiroElasticsearch struct {
	Con *sql.DB
}

func (e *EmiroElasticsearch) SearchQuery(query string, count int32, all bool) (*sql.Rows, error) {
	return nil, nil
}

func (e *EmiroElasticsearch) CreateOrConnect() error {
	log.Println("CREATE OR CONNECT ELASTIC")
	es, _ := elasticsearch.NewDefaultClient()
	log.Println(elasticsearch.Version)
	log.Println(es.Info())

	// cfg := elasticsearch.Config{
	// 	Addresses: []string{
	// 		"http://localhost:9200",
	// 	},
	// }
	// es, err := elasticsearch.NewClient(cfg)
	return nil

}
