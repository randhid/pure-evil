package evilinput

import (
	"context"
	"errors"
	"evil-modules/common"
	"sync"
	"time"

	"go.viam.com/rdk/components/input"
	"go.viam.com/rdk/logging"
	"go.viam.com/rdk/resource"
)

const inputName = "evil-input"

var Model = common.EvilsFamily.WithModel(inputName)

func init() {
	resource.RegisterComponent(input.API, Model, resource.Registration[input.Controller, resource.NoNativeConfig]{
		Constructor: newevilInput,
	})
}

func newevilInput(ctx context.Context, deps resource.Dependencies, conf resource.Config, logger logging.Logger) (
	input.Controller, error,
) {
	return &evil{Named: conf.ResourceName().AsNamed()}, nil
}

// An InputController evils an input.Controller.
type evil struct {
	resource.Named
	resource.TriviallyCloseable
	resource.TriviallyReconfigurable
	mu sync.Mutex
}

// Controls lists the inputs of the gamepad.
func (f *evil) Controls(ctx context.Context, extra map[string]interface{}) ([]input.Control, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	time.Sleep(time.Minute)
	return nil, nil
}

// Events returns the a specified or random input.Event (the current state) for AbsoluteX.
func (f *evil) Events(ctx context.Context, extra map[string]interface{}) (map[input.Control]input.Event, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	time.Sleep(time.Minute)
	return nil, nil
}

// RegisterControlCallback registers a callback function to be executed on the specified trigger Event. The evil implementation will
// trigger the callback at a random or user-specified interval with a random or user-specified value.
func (f *evil) RegisterControlCallback(
	ctx context.Context,
	control input.Control,
	triggers []input.EventType,
	ctrlFunc input.ControlFunction,
	extra map[string]interface{},
) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	time.Sleep(time.Minute)
	return nil
}

// TriggerEvent allows directly sending an Event (such as a button press) from external code.
func (f *evil) TriggerEvent(ctx context.Context, event input.Event, extra map[string]interface{}) error {
	return errors.New("unsupported")
}
