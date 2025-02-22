package source

import (
	"fmt"

	"github.com/livemart/psql-streamer/common"
	"github.com/livemart/psql-streamer/sink"
	kafkasrc "github.com/livemart/psql-streamer/source/kafka"
	"github.com/livemart/psql-streamer/source/postgres"
	"github.com/spf13/viper"
)

// Source represent a generic source that produces events
type Source interface {
	Subscribe(sink.Sink)
	Start()
	Flush() error

	common.Common
}

// Init initializes a source from a viper subtree
func Init(name string, v *viper.Viper) (s Source, err error) {
	t := v.GetString("type")
	if t == "" {
		return nil, fmt.Errorf("Sink type not specified")
	}

	switch t {
	case "kafka":
		return kafkasrc.New(name, v)
	case "postgres":
		return postgres.New(name, v)
	default:
		return nil, fmt.Errorf("Unknown sink type: %s", t)
	}
}
