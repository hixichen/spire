package telemetry

import (
	"context"
	"github.com/andres-erbsen/clock"
	"time"
)

func EmitUptime(ctx context.Context, itl time.Duration, m Metrics) error {
	clk := clock.New()
	startTime := clk.Now()

	if len(itl.String()) == 0 {
		// by default, report every 10 seconds.
		itl = time.Second * 10
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-clk.After(itl):
				m.SetGauge([]string{"uptime"}, float32(clk.Now().Sub(startTime)/time.Millisecond))
			}
		}
	}()

	return nil
}
