package evilmovementsensor

import (
	"context"
	"math"
	"sync"
	"time"

	"evil-modules/common"

	"github.com/golang/geo/r3"
	geo "github.com/kellydunn/golang-geo"
	"go.viam.com/rdk/components/movementsensor"
	"go.viam.com/rdk/grpc"
	"go.viam.com/rdk/logging"
	"go.viam.com/rdk/resource"
	"go.viam.com/rdk/spatialmath"
)

const (
	movementsensorName = "evil-movementsensor"
)

var (
	Model = common.EvilsFamily.WithModel(movementsensorName)
	nan   = math.NaN()
)

func init() {
	resource.RegisterComponent(movementsensor.API, Model, resource.Registration[movementsensor.MovementSensor, resource.NoNativeConfig]{
		Constructor: newevilMovementSensor,
	})
}

type evil struct {
	resource.Named
	resource.TriviallyReconfigurable
	resource.TriviallyCloseable
	resource.Sensor
	mu sync.Mutex
}

func newevilMovementSensor(ctx context.Context, deps resource.Dependencies, conf resource.Config, logger logging.Logger) (
	movementsensor.MovementSensor, error,
) {
	f := &evil{
		Named: conf.ResourceName().AsNamed(),
	}

	if err := f.Reconfigure(ctx, deps, conf); err != nil {
		return nil, err
	}

	return f, nil
}

func (f *evil) Reconfigure(ctx context.Context, deps resource.Dependencies, conf resource.Config) error {
	_, err := resource.NativeConfig[resource.NoNativeConfig](conf)
	if err != nil {
		return err
	}

	f.mu.Lock()
	defer f.mu.Unlock()
	time.Sleep(time.Minute)

	return nil
}

func (f *evil) Position(ctx context.Context, extra map[string]interface{}) (*geo.Point, float64, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	time.Sleep(time.Minute)

	newcoord := geo.NewPoint(nan, nan)
	return newcoord, nan, nil
}

func (f *evil) Orientation(ctx context.Context, extra map[string]interface{}) (spatialmath.Orientation, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	time.Sleep(time.Minute)

	return &spatialmath.OrientationVector{OX: nan, OY: nan, OZ: nan, Theta: nan}, nil
}

func (f *evil) CompassHeading(ctx context.Context, extra map[string]interface{}) (float64, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	time.Sleep(time.Minute)
	return nan, grpc.UnimplementedError
}

func (f *evil) AngularVelocity(ctx context.Context, extra map[string]interface{}) (spatialmath.AngularVelocity, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	time.Sleep(time.Minute)

	return spatialmath.AngularVelocity{X: nan, Y: nan, Z: nan}, nil
}

func (f *evil) LinearVelocity(ctx context.Context, extra map[string]interface{}) (r3.Vector, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	time.Sleep(time.Minute)
	return r3.Vector{X: nan, Y: nan, Z: nan}, nil
}

func (f *evil) LinearAcceleration(ctx context.Context, extra map[string]interface{}) (r3.Vector, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	time.Sleep(time.Minute)

	return r3.Vector{X: nan, Y: nan, Z: nan}, nil
}

func (f *evil) Accuracy(ctx context.Context, extra map[string]interface{}) (*movementsensor.Accuracy, error) {
	return &movementsensor.Accuracy{
		AccuracyMap:        nil,
		Hdop:               float32(nan),
		Vdop:               float32(nan),
		CompassDegreeError: float32(nan),
		NmeaFix:            int32(nan),
	}, nil
}

func (f *evil) Properties(ctx context.Context, extra map[string]interface{}) (*movementsensor.Properties, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	return &movementsensor.Properties{
		PositionSupported:           true,
		OrientationSupported:        true,
		CompassHeadingSupported:     true,
		LinearVelocitySupported:     true,
		AngularVelocitySupported:    true,
		LinearAccelerationSupported: true,
	}, nil
}
