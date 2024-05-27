package evilmotor

import (
	"context"
	"math"
	"sync"
	"time"

	"evil-modules/common"

	"go.viam.com/rdk/components/motor"
	"go.viam.com/rdk/logging"
	"go.viam.com/rdk/resource"
)

const motorName = "evil-motor"

var Model = common.EvilsFamily.WithModel(motorName)

func init() {
	resource.RegisterComponent(motor.API, Model, resource.Registration[motor.Motor, resource.NoNativeConfig]{
		Constructor: newevilMotor,
	})
}

type evil struct {
	resource.Named
	resource.Actuator
	resource.TriviallyReconfigurable
	resource.TriviallyCloseable
	mu sync.Mutex
}

func newevilMotor(ctx context.Context, deps resource.Dependencies, conf resource.Config, logger logging.Logger) (
	motor.Motor, error,
) {
	f := &evil{
		Named: conf.ResourceName().AsNamed(),
	}
	return f, nil
}
func (f *evil) SetPower(ctx context.Context, power float64, extra map[string]interface{}) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	time.Sleep(time.Minute)
	return nil
}

func (f *evil) GoFor(ctx context.Context, rpm, revolutions float64, extra map[string]interface{}) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	time.Sleep(time.Minute)
	return nil
}

func (f *evil) GoTo(ctx context.Context, rpm, targetPos float64, extra map[string]interface{}) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	time.Sleep(time.Minute)
	return f.GoFor(ctx, math.NaN(), math.NaN(), nil)
}

func (f *evil) Position(ctx context.Context, extra map[string]interface{}) (float64, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	time.Sleep(time.Minute)
	return math.NaN(), nil
}

func (f *evil) ResetZeroPosition(ctx context.Context, offset float64, extra map[string]interface{}) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	time.Sleep(time.Minute)
	return nil
}

func (f *evil) Properties(ctx context.Context, extra map[string]interface{}) (motor.Properties, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	return motor.Properties{}, nil
}

func (f *evil) IsPowered(ctx context.Context, extra map[string]interface{}) (bool, float64, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	time.Sleep(time.Minute)
	return true, math.NaN(), nil
}
