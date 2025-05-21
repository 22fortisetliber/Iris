package distributor

import (
	"fmt"
	"time"

	"github.com/go-kit/log"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	"github.com/22fortisetliber/iris/pkg/ring"
	ring_client "github.com/22fortisetliber/iris/pkg/ring/client"
)

var clients = promauto.NewGauge(prometheus.GaugeOpts{
	Namespace: "iris",
	Name:      "dispatcher_clients",
	Help:      "Number of dispatcher clients",
})

type PoolConfig struct {
	ClientCleanUpInterval time.Duration `yaml:"client_cleanup_interval"`
	HealthCheckDispatcher bool          `yaml:"health_check_dispatcher"`
	RemoteTimeout         time.Duration `yaml:"remote_timeout"`
}

func (c *PoolConfig) Validate() error {
	if c.ClientCleanUpInterval <= 0 {
		return fmt.Errorf("client_cleanup_interval must be greater than 0")
	}
	if c.RemoteTimeout <= 0 {
		return fmt.Errorf("remote_timeout must be greater than 0")
	}
	return nil
}

func NewPool(cfg PoolConfig, ring ring.ReadRing, factory ring_client.PoolFactory, logger log.Logger) *ring_client.Pool {
	poolCfg := ring_client.PoolConfig{
		CheckInterval:      cfg.ClientCleanUpInterval,
		HealthCheckEnabled: cfg.HealthCheckDispatcher,
		HealthCheckTimeout: cfg.RemoteTimeout,
	}
	return ring_client.NewPool("dispatcher", poolCfg, ring_client.NewRingServiceDiscovery(ring), factory, clients, logger)
}