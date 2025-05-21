package distributor

import (
	"time"

	"github.com/go-kit/log"
	"github.com/prometheus/client_golang/prometheus"
	// "github.com/go-kit/log/level"

	"github.com/22fortisetliber/iris/pkg/ring"
	ring_client "github.com/22fortisetliber/iris/pkg/ring/client"
	"github.com/22fortisetliber/iris/pkg/util/services"
)

type Distributor struct {
	services.Service

	cfg            Config
	log            log.Logger
	dispatcherRing ring.ReadRing
	dispatcherPool *ring_client.Pool

	distributorLifeCycler *ring.Lifecycler
	distributorRing       *ring.Ring

	// Manager for subservices (HA Tracker, distributor ring and client pool)
	subservices        *services.Manager
	subservicesWatcher *services.FailureWatcher

	// Metrics
	receivedAlerts prometheus.CounterVec
}

type Config struct {
	PoolConfig       PoolConfig    `yaml:"pool_config"`
	RemoteTimeout    time.Duration `yaml:"remote_timeout"`
	ShardingStrategy string        `yaml:"sharding_strategy"`
}
