package config

import (
	"fmt"
	"errors"
)

var (
	ErrInvalidType = errors.New("invalid type entry in config")
)

type validatorContext struct {
	usedPaths map[string]bool
}

var validators = map[string]func(*validatorContext, map[string]interface{}) error{}

func init() {
	validators["badgerds"] = badgerdsValidator
	validators["flatfs"] = flatfsValidator
	validators["levelds"] = leveldsValidator
	validators["measure"] = measureValidator
	validators["mount"] = mountValidator
}

func Validate(dsConfiguration map[string]interface{}) error {
	ctx := validatorContext{
		usedPaths: map[string]bool{},
	}
	return validate(&ctx, dsConfiguration)
}

func validate(ctx *validatorContext, dsConfiguration map[string]interface{}) error {
	t, ok := dsConfiguration["type"].(string)
	if !ok {
		return ErrInvalidType
	}

	validator := validators[t]
	if validator == nil {
		return fmt.Errorf("unsupported type entry in config: %s", t)
	}

	return validator(ctx, dsConfiguration)
}

func checkPath(ctx *validatorContext, p interface{}) error {
	path, ok := p.(string)
	if !ok {
		return errors.New("invalid 'path' type in flatfs datastore")
	}

	if ctx.usedPaths[path] {
		return fmt.Errorf("Path '%s' is already in use", path)
	}

	ctx.usedPaths[path] = true

	//TODO: better path validation

	return nil
}

//////////////

func flatfsValidator(ctx *validatorContext, dsConfiguration map[string]interface{}) error {
	err := checkPath(ctx, dsConfiguration["path"])
	if err != nil {
		return err
	}

	return nil
}

func leveldsValidator(ctx *validatorContext, dsConfiguration map[string]interface{}) error {
	err := checkPath(ctx, dsConfiguration["path"])
	if err != nil {
		return err
	}

	_, ok := dsConfiguration["compression"].(string)
	if !ok {
		return errors.New("invalid 'compression' type in levelds datastore")
	}

	return nil
}

func badgerdsValidator(ctx *validatorContext, dsConfiguration map[string]interface{}) error {
	err := checkPath(ctx, dsConfiguration["path"])
	if err != nil {
		return err
	}

	return nil
}

func mountValidator(ctx *validatorContext, dsConfiguration map[string]interface{}) error {
	mounts, ok := dsConfiguration["mounts"].([]interface{})
	if !ok {
		return errors.New("invalid 'mounts' in mount datastore")
	}

	mountPoints := map[string]bool{}

	for _, m := range mounts {
		mount, ok := m.(map[string]interface{})
		if !ok {
			return errors.New("mounts entry has invalid type")
		}

		mountPoint, ok := mount["mountpoint"].(string)
		if !ok {
			return errors.New("'mountpoint' must be a string")
		}

		if mountPoints[mountPoint] {
			return errors.New("multiple mounts under one path are not allowed")
		}

		mountPoints[mountPoint] = true

		err := validate(ctx, mount)
		if err != nil {
			return err
		}
	}

	return nil
}

func measureValidator(ctx *validatorContext, dsConfiguration map[string]interface{}) error {
	_, ok := dsConfiguration["prefix"].(string)
	if !ok {
		return errors.New("invalid 'prefix' in measure datastore")
	}

	child, ok := dsConfiguration["child"].(map[string]interface{})
	if !ok {
		return errors.New("child of measure datastore has invalid type")
	}

	return validate(ctx, child)
}