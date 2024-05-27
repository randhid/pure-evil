package evilservo

import (
	"context"
	"evil-modules/common"
	"math"
	"sync"
	"time"

	"go.viam.com/rdk/components/servo"
	"go.viam.com/rdk/logging"
	"go.viam.com/rdk/resource"
)

const (
	servoName = "evil-servo"
	increment = 1
)

var Model = common.EvilsFamily.WithModel(servoName)

func init() {
	resource.RegisterComponent(servo.API, Model, resource.Registration[servo.Servo, resource.NoNativeConfig]{
		Constructor: newevilServo,
	})
}

type evil struct {
	resource.Named
	resource.TriviallyReconfigurable
	resource.TriviallyCloseable
	resource.Actuator

	mu sync.Mutex
}

func newevilServo(ctx context.Context, deps resource.Dependencies, conf resource.Config, logger logging.Logger) (
	servo.Servo, error,
) {
	f := &evil{
		Named: conf.ResourceName().AsNamed(),
	}

	return f, nil
}

func (f *evil) Move(ctx context.Context, pos uint32, extra map[string]interface{}) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	time.Sleep(time.Minute)
	return nil
}

func (f *evil) Position(ctx context.Context, extra map[string]interface{}) (uint32, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	time.Sleep(time.Minute)
	return uint32(math.NaN()), nil
}
