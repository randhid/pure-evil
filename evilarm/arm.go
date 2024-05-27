package evilarm

import (
	// for embedding model file.
	"context"
	_ "embed"
	"evil-modules/common"
	"sync"
	"time"

	pb "go.viam.com/api/component/arm/v1"
	"go.viam.com/rdk/components/arm"
	"go.viam.com/rdk/grpc"
	"go.viam.com/rdk/logging"
	"go.viam.com/rdk/referenceframe"
	"go.viam.com/rdk/resource"
	"go.viam.com/rdk/spatialmath"
)

const armName = "evil-arm"

var Model = common.EvilsFamily.WithModel(armName)

//go:embed kinematics.json
var kinematics []byte

func init() {
	resource.RegisterComponent(arm.API, Model, resource.Registration[arm.Arm, resource.NoNativeConfig]{
		Constructor: newevilArm,
	})
}

type evil struct {
	resource.Named
	resource.TriviallyReconfigurable
	resource.TriviallyCloseable
	referenceframe.InputEnabled
	referenceframe.ModelFramer
	resource.Actuator
	resource.Shaped

	mu sync.Mutex
}

func newevilArm(ctx context.Context, deps resource.Dependencies, conf resource.Config, logger logging.Logger) (
	arm.Arm, error,
) {
	f := &evil{
		Named: conf.ResourceName().AsNamed(),
	}
	return f, nil
}

// MoveToJointPositions sets the joints.
func (f *evil) MoveToJointPositions(ctx context.Context, joints *pb.JointPositions, extra map[string]interface{}) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	time.Sleep(1 * time.Minute)

	return nil
}

// JointPositions returns joints.
func (f *evil) JointPositions(ctx context.Context, extra map[string]interface{}) (*pb.JointPositions, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	time.Sleep(1 * time.Minute)
	return nil, nil
}

func (f *evil) EndPosition(ctx context.Context, extra map[string]interface{}) (spatialmath.Pose, error) {
	return spatialmath.NewZeroPose(), grpc.UnimplementedError
}

func (f *evil) MoveToPosition(ctx context.Context, pos spatialmath.Pose, extra map[string]interface{}) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	time.Sleep(1 * time.Minute)
	return grpc.UnimplementedError
}
