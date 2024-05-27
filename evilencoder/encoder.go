package evilencoder

import (
	"context"
	"evil-modules/common"
	"math"
	"sync"
	"time"

	"go.viam.com/rdk/components/encoder"
	"go.viam.com/rdk/logging"
	"go.viam.com/rdk/resource"
)

const (
	encoderName  = "evil-encoder"
	setIncrement = "setIncrement" // for DoCommand to make the increment larger or siwtch directions
)

var (
	Model = common.EvilsFamily.WithModel(encoderName)
)

func init() {
	resource.RegisterComponent(encoder.API, Model, resource.Registration[encoder.Encoder, resource.NoNativeConfig]{
		Constructor: newevilEncoder,
	})
}

type evil struct {
	resource.TriviallyReconfigurable
	resource.TriviallyCloseable
	resource.Named

	logger logging.Logger

	mu sync.Mutex
}

func newevilEncoder(ctx context.Context, deps resource.Dependencies, conf resource.Config, logger logging.Logger) (
	encoder.Encoder, error,
) {
	f := &evil{
		Named: conf.ResourceName().AsNamed(),
	}

	return f, nil
}

func (f *evil) Position(ctx context.Context, encType encoder.PositionType, extra map[string]interface{}) (float64, encoder.PositionType, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	return math.NaN(), encoder.PositionTypeUnspecified, nil
}

func (f *evil) ResetPosition(ctx context.Context, extra map[string]interface{}) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	time.Sleep(time.Minute)
	return nil
}

func (f *evil) Properties(ctx context.Context, extra map[string]interface{}) (encoder.Properties, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	time.Sleep(time.Minute)
	return encoder.Properties{AngleDegreesSupported: true, TicksCountSupported: true}, nil
}
