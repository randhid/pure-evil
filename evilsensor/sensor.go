package evilsensor

import (
	"context"
	"evil-modules/common"
	"sync"

	"go.viam.com/rdk/components/sensor"
	"go.viam.com/rdk/logging"
	"go.viam.com/rdk/resource"
)

const sensorName = "evil-sensor"

var Model = common.EvilsFamily.WithModel(sensorName)

type Config struct {
	resource.TriviallyValidateConfig
}

func init() {
	resource.RegisterComponent(sensor.API, Model, resource.Registration[sensor.Sensor, resource.NoNativeConfig]{
		Constructor: newevilSensor,
	})
}

type evil struct {
	resource.Named
	resource.TriviallyReconfigurable
	resource.TriviallyCloseable
	resource.Sensor

	mu sync.Mutex
}

func newevilSensor(ctx context.Context, deps resource.Dependencies, conf resource.Config, logger logging.Logger) (
	sensor.Sensor, error,
) {
	f := &evil{
		Named: conf.ResourceName().AsNamed(),
	}

	return f, nil
}
