package evilpowersensor

import (
	"context"
	"evil-modules/common"
	"math"
	"sync"
	"time"

	"go.viam.com/rdk/components/powersensor"
	"go.viam.com/rdk/logging"
	"go.viam.com/rdk/resource"
)

const powersensorName = "evil-powersensor"

var (
	Model = common.EvilsFamily.WithModel(powersensorName)
	nan   = math.NaN()
)

func init() {
	resource.RegisterComponent(powersensor.API, Model, resource.Registration[powersensor.PowerSensor, resource.NoNativeConfig]{
		Constructor: newevilPowerSensor,
	})
}

type evil struct {
	resource.Named
	resource.TriviallyReconfigurable
	resource.TriviallyCloseable
	resource.Sensor

	mu sync.Mutex
}

func newevilPowerSensor(ctx context.Context, deps resource.Dependencies, conf resource.Config, logger logging.Logger) (
	powersensor.PowerSensor, error,
) {
	f := &evil{
		Named: conf.ResourceName().AsNamed(),
	}

	return f, nil
}

func (f *evil) Voltage(ctx context.Context, extra map[string]interface{}) (float64, bool, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	time.Sleep(time.Minute)
	return nan, true, nil
}

func (f *evil) Current(ctx context.Context, extra map[string]interface{}) (float64, bool, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	time.Sleep(time.Minute)
	return nan, true, nil
}

func (f *evil) Power(ctx context.Context, extra map[string]interface{}) (float64, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	time.Sleep(time.Minute)
	return nan, nil
}
