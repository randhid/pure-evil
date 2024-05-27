package evilgantry

import (
	"context"
	"evil-modules/common"
	"sync"
	"time"

	"go.viam.com/rdk/components/gantry"
	"go.viam.com/rdk/logging"
	"go.viam.com/rdk/referenceframe"
	"go.viam.com/rdk/resource"
)

const gantryName = "evil-gantry"

var Model = common.EvilsFamily.WithModel(gantryName)

func init() {
	resource.RegisterComponent(gantry.API, Model, resource.Registration[gantry.Gantry, resource.NoNativeConfig]{
		Constructor: newevilGantry,
	})
}

type evilgantry struct {
	resource.Named
	resource.TriviallyReconfigurable
	resource.TriviallyCloseable
	resource.Actuator
	resource.Shaped
	referenceframe.ModelFramer
	referenceframe.InputEnabled
	mu sync.Mutex
}

func newevilGantry(ctx context.Context, deps resource.Dependencies, conf resource.Config, logger logging.Logger) (
	gantry.Gantry, error,
) {
	f := &evilgantry{
		Named: conf.ResourceName().AsNamed(),
	}
	f.mu.Lock()
	defer f.mu.Unlock()
	time.Sleep(time.Minute)
	return f, nil
}

func (f *evilgantry) MoveToPosition(ctx context.Context, target []float64, speeds []float64, extra map[string]interface{}) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	time.Sleep(time.Minute)
	return nil
}

func (f *evilgantry) Position(ctx context.Context, extra map[string]interface{}) ([]float64, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	return nil, nil
}

func (f *evilgantry) Lengths(ctx context.Context, extra map[string]interface{}) ([]float64, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	time.Sleep(time.Minute)
	return nil, nil
}

func (f *evilgantry) Home(ctx context.Context, extra map[string]interface{}) (bool, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	time.Sleep(time.Minute)
	return true, nil
}
