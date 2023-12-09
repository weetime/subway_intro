package data

import (
	"strings"

	"subway_intro/internal/conf"
	"subway_intro/internal/data/ent"
	"github.com/go-kratos/kratos/v2/log"

	"github.com/google/wire"
	_ "github.com/lib/pq"
)

var ProviderSet = wire.NewSet(
	NewData,
	NewUserRepo,
)

// Data .
type Data struct {
	db *ent.Client
}

// NewData .和ent放到同级目录，可以支持多库
func NewData(conf *conf.Bootstrap, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(log.With(logger, "module", "catalog-service/data"))

	var options []ent.Option

	if conf.Data.Database.Debug {
		options = append(options, ent.Debug())
	}

	client, err := ent.Open(
		strings.ToLower(conf.Data.Database.Driver),
		conf.Data.Database.Source,
		options...,
	)

	if err != nil {
		log.Errorf("failed opening connection to postgres: %v", err)
		return nil, nil, err
	}

	d := &Data{
		db: client,
	}

	return d, func() {
		log.Info("message", "closing the data resources")
		if err := client.Close(); err != nil {
			log.Error(err)
		}
	}, nil
}
