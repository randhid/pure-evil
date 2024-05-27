package evilgripper

import (
	"context"
	"evil-modules/common"
	"sync"
	"time"

	"go.viam.com/rdk/components/gripper"
	"go.viam.com/rdk/logging"
	"go.viam.com/rdk/referenceframe"
	"go.viam.com/rdk/resource"
)

const (
	servoName = "evil-gripper"
	increment = 1
)

var Model = common.EvilsFamily.WithModel(servoName)

func init() {
	resource.RegisterComponent(gripper.API, Model, resource.Registration[gripper.Gripper, resource.NoNativeConfig]{
		Constructor: newevilGripper,
	})
}

type evil struct {
	resource.Named
	resource.TriviallyReconfigurable
	resource.TriviallyCloseable
	resource.Shaped
	resource.Actuator
	referenceframe.ModelFramer
	mu sync.Mutex
}

func newevilGripper(ctx context.Context, deps resource.Dependencies, conf resource.Config, logger logging.Logger) (
	gripper.Gripper, error,
) {
	f := &evil{
		Named: conf.ResourceName().AsNamed(),
	}

	return f, nil
}

func (f *evil) Grab(ctx context.Context, extra map[string]interface{}) (bool, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	time.Sleep(time.Minute)
	return true, nil
}

func (f *evil) Open(ctx context.Context, extra map[string]interface{}) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	time.Sleep(time.Minute)
	return nil
}
