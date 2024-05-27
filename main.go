package main

import (
	"context"
	"evil-modules/evilarm"
	"evil-modules/evilbase"
	"evil-modules/evilboard"
	"evil-modules/evilcamera"
	"evil-modules/evilencoder"
	"evil-modules/evilgantry"
	"evil-modules/evilgripper"
	"evil-modules/evilinput"
	"evil-modules/evilmotor"
	"evil-modules/evilmovementsensor"
	"evil-modules/evilpowersensor"
	"evil-modules/evilsensor"
	"evil-modules/evilservo"

	"go.viam.com/rdk/components/arm"
	"go.viam.com/rdk/components/base"
	"go.viam.com/rdk/components/board"
	"go.viam.com/rdk/components/camera"
	"go.viam.com/rdk/components/encoder"
	"go.viam.com/rdk/components/gantry"
	"go.viam.com/rdk/components/gripper"
	"go.viam.com/rdk/components/input"
	"go.viam.com/rdk/components/motor"
	"go.viam.com/rdk/components/movementsensor"
	"go.viam.com/rdk/components/powersensor"
	"go.viam.com/rdk/components/sensor"
	"go.viam.com/rdk/components/servo"
	"go.viam.com/rdk/logging"
	"go.viam.com/rdk/module"
	"go.viam.com/rdk/resource"
	"go.viam.com/utils"
)

var ModuleFamily = resource.NewModelFamily("rand", "go-evils")

func main() {
	utils.ContextualMain(mainWithArgs, module.NewLoggerFromArgs("evil Go Modules"))
}

func mainWithArgs(ctx context.Context, args []string, logger logging.Logger) (err error) {
	// instantiates the module itself
	evils, err := module.NewModuleFromArgs(ctx, logger)
	if err != nil {
		return err
	}

	// Models and APIs add helpers to the registry during their init().
	// They can then be added to the module here.s
	if err = evils.AddModelFromRegistry(ctx, arm.API, evilarm.Model); err != nil {
		return err
	}

	if err = evils.AddModelFromRegistry(ctx, base.API, evilbase.Model); err != nil {
		return err
	}

	if err = evils.AddModelFromRegistry(ctx, board.API, evilboard.Model); err != nil {
		return err
	}

	if err = evils.AddModelFromRegistry(ctx, camera.API, evilcamera.Model); err != nil {
		return err
	}

	if err = evils.AddModelFromRegistry(ctx, encoder.API, evilencoder.Model); err != nil {
		return err
	}

	if err = evils.AddModelFromRegistry(ctx, gantry.API, evilgantry.Model); err != nil {
		return err
	}

	if err = evils.AddModelFromRegistry(ctx, gripper.API, evilgripper.Model); err != nil {
		return err
	}

	if err = evils.AddModelFromRegistry(ctx, input.API, evilinput.Model); err != nil {
		return err
	}

	if err = evils.AddModelFromRegistry(ctx, motor.API, evilmotor.Model); err != nil {
		return err
	}

	if err = evils.AddModelFromRegistry(ctx, movementsensor.API, evilmovementsensor.Model); err != nil {
		return err
	}

	if err = evils.AddModelFromRegistry(ctx, powersensor.API, evilpowersensor.Model); err != nil {
		return err
	}

	if err = evils.AddModelFromRegistry(ctx, servo.API, evilservo.Model); err != nil {
		return err
	}

	if err = evils.AddModelFromRegistry(ctx, sensor.API, evilsensor.Model); err != nil {
		return err
	}

	// Each module runs as its own process
	err = evils.Start(ctx)
	logger.Warn("starting module")
	defer evils.Close(ctx)
	if err != nil {
		return err
	}
	<-ctx.Done()
	return nil
}
