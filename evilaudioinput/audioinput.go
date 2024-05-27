//go:build !no_cgo

// Package evil implements a evil audio input.
package evil

import (
	"context"
	"evil-modules/common"
	"sync"
	"time"

	"github.com/pion/mediadevices/pkg/prop"
	"github.com/pion/mediadevices/pkg/wave"

	"go.viam.com/rdk/components/audioinput"
	"go.viam.com/rdk/gostream"
	"go.viam.com/rdk/grpc"
	"go.viam.com/rdk/logging"
	"go.viam.com/rdk/resource"
)

const armName = "evil-audio-in"

var Model = common.EvilsFamily.WithModel(armName)

func init() {
	resource.RegisterComponent(audioinput.API, Model, resource.Registration[audioinput.AudioInput, resource.NoNativeConfig]{
		Constructor: newevilAudio,
	})
}

func newevilAudio(ctx context.Context, deps resource.Dependencies, conf resource.Config, logger logging.Logger) (
	audioinput.AudioInput, error,
) {

	i := &evil{
		Named:  conf.ResourceName().AsNamed(),
		logger: logger,
	}

	return i, nil

}

// audioInput is a evil audioinput that always returns the same chunk.
type evil struct {
	resource.Named
	resource.TriviallyReconfigurable
	resource.TriviallyCloseable
	logger logging.Logger
	mu     sync.Mutex
}

const (
	latencyMillis = 20
	samplingRate  = 48000
	channelCount  = 1
)

func (f *evil) Read(ctx context.Context) (wave.Audio, func(), error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	time.Sleep(time.Minute)
	return nil, func() {}, grpc.UnimplementedError
}

func (f *evil) MediaProperties(_ context.Context) (prop.Audio, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	time.Sleep(time.Minute)
	return prop.Audio{}, grpc.UnimplementedError
}

func (f *evil) Stream(context.Context, ...gostream.ErrorHandler) (gostream.MediaStream[wave.Audio], error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	time.Sleep(time.Minute)
	return nil, grpc.UnimplementedError
}
