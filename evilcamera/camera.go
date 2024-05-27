package evilcamera

import (
	"context"
	"evil-modules/common"
	"image"
	"time"

	"sync"

	"go.viam.com/rdk/components/camera"
	"go.viam.com/rdk/gostream"
	"go.viam.com/rdk/grpc"
	"go.viam.com/rdk/logging"
	"go.viam.com/rdk/pointcloud"
	"go.viam.com/rdk/resource"
	"go.viam.com/rdk/rimage/transform"
)

const cameraName = "evil-camera"

var Model = common.EvilsFamily.WithModel(cameraName)

func init() {
	resource.RegisterComponent(camera.API, Model, resource.Registration[camera.Camera, resource.NoNativeConfig]{
		Constructor: newevilCamera,
	})
}

type evil struct {
	resource.Named
	resource.TriviallyReconfigurable
	resource.TriviallyCloseable
	logger logging.Logger

	mu sync.Mutex
}

func newevilCamera(ctx context.Context, deps resource.Dependencies, conf resource.Config, logger logging.Logger) (
	camera.Camera, error,
) {
	f := &evil{
		Named: conf.ResourceName().AsNamed(),
	}

	return f, nil
}

func (f *evil) Images(context.Context) ([]camera.NamedImage, resource.ResponseMetadata, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	time.Sleep(time.Minute)
	return []camera.NamedImage{}, resource.ResponseMetadata{}, grpc.UnimplementedError
}

func (f *evil) NextPointCloud(context.Context) (pointcloud.PointCloud, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	time.Sleep(time.Minute)
	return pointcloud.New(), grpc.UnimplementedError
}

func (f *evil) Projector(context.Context) (transform.Projector, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	time.Sleep(time.Minute)
	return nil, grpc.UnimplementedError
}

func (f *evil) Properties(context.Context) (camera.Properties, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	time.Sleep(time.Minute)
	return camera.Properties{}, grpc.UnimplementedError
}

func (f *evil) Stream(context.Context, ...gostream.ErrorHandler) (gostream.MediaStream[image.Image], error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	time.Sleep(time.Minute)
	return nil, grpc.UnimplementedError
}
