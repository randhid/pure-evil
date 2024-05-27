package evilboard

import (
	"context"
	"evil-modules/common"
	"sync"
	"time"

	pb "go.viam.com/api/component/board/v1"
	"go.viam.com/rdk/components/board"
	"go.viam.com/rdk/grpc"
	"go.viam.com/rdk/logging"
	"go.viam.com/rdk/resource"
)

const boardName = "evil-board"

var Model = common.EvilsFamily.WithModel(boardName)

func init() {
	resource.RegisterComponent(board.API, Model, resource.Registration[board.Board, *resource.NoNativeConfig]{
		Constructor: newevilBoard,
	})
}

type evilboard struct {
	resource.Named
	resource.TriviallyReconfigurable

	mu sync.Mutex
}

func newevilBoard(ctx context.Context, deps resource.Dependencies, conf resource.Config, logger logging.Logger) (
	board.Board, error,
) {
	f := &evilboard{
		Named: conf.ResourceName().AsNamed(),
	}

	return f, nil
}

// AnalogByName returns the analog pin by the given name if it exists.
func (f *evilboard) AnalogByName(name string) (board.Analog, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	time.Sleep(time.Minute)
	return nil, nil
}

// DigitalInterruptByName returns the interrupt by the given name if it exists.
func (f *evilboard) DigitalInterruptByName(name string) (board.DigitalInterrupt, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	time.Sleep(time.Minute)
	return nil, nil
}

// GPIOPinByName returns the GPIO pin by the given name if it exists.
func (f *evilboard) GPIOPinByName(name string) (board.GPIOPin, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	time.Sleep(time.Minute)
	return nil, nil
}

// AnalogNames returns the names of all known analog pins.
func (f *evilboard) AnalogNames() []string {
	return nil
}

// DigitalInterruptNames returns the names of all known digital interrupts.
func (f *evilboard) DigitalInterruptNames() []string {
	return nil
}

// SetPowerMode sets the board to the given power mode. If provided,
// the board will exit the given power mode after the specified
// duration.
func (f *evilboard) SetPowerMode(ctx context.Context, mode pb.PowerMode, duration *time.Duration) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	time.Sleep(time.Minute)
	return grpc.UnimplementedError
}

// StreamTicks starts a stream of digital interrupt ticks.
func (f *evilboard) StreamTicks(ctx context.Context, interrupts []board.DigitalInterrupt, ch chan board.Tick,
	extra map[string]interface{},
) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	time.Sleep(time.Minute)
	return nil
}

// Close attempts to cleanly close each part of the board.
func (f *evilboard) Close(ctx context.Context) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	time.Sleep(time.Minute)
	return nil
}
