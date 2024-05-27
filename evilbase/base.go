package evilbase

import (
	"context"
	"evil-modules/common"
	"sync"
	"time"

	"github.com/golang/geo/r3"
	"go.viam.com/rdk/components/base"
	"go.viam.com/rdk/logging"
	"go.viam.com/rdk/resource"
)

const (
	baseName = "evil-base"
)

var Model = common.EvilsFamily.WithModel(baseName)

func init() {
	resource.RegisterComponent(base.API, Model, resource.Registration[base.Base, resource.NoNativeConfig]{
		Constructor: newevilBase,
	})
}

type evil struct {
	resource.Named
	resource.TriviallyReconfigurable
	resource.TriviallyCloseable
	resource.Shaped
	resource.Actuator
	logger logging.Logger

	mu sync.Mutex
}

func newevilBase(ctx context.Context, deps resource.Dependencies, conf resource.Config, logger logging.Logger) (
	base.Base, error,
) {
	f := &evil{
		Named: conf.ResourceName().AsNamed(),
	}

	return f, nil
}

func (f *evil) SetPower(ctx context.Context, linear, angular r3.Vector, extra map[string]interface{}) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	time.Sleep(time.Minute)
	return nil
}

func (f *evil) MoveStraight(ctx context.Context, distance int, mmPerSec float64, extra map[string]interface{}) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	time.Sleep(time.Minute)
	return nil
}

func (f *evil) Spin(ctx context.Context, angle, degsPerSec float64, extra map[string]interface{}) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	time.Sleep(time.Minute)
	return nil
}

func (f *evil) SetVelocity(ctx context.Context, linear, angular r3.Vector, extra map[string]interface{}) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	time.Sleep(time.Minute)
	return nil
}

func (f *evil) Close(ctx context.Context) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	time.Sleep(time.Minute)
	return f.Stop(ctx, nil)
}

func (f *evil) Properties(ctx context.Context, extra map[string]interface{}) (base.Properties, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	time.Sleep(time.Minute)
	return base.Properties{}, nil
}
