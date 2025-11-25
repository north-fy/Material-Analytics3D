package calculator

import (
	"errors"
)

// errors
var (
	ErrEnoughArg   error = errors.New("not enough arguments for calculations")
	ErrUnknownType error = errors.New("unknown object type for calculations")
)

// math
const (
	CubeType = iota
	ParallelepipedType
	PyramidType
	//ConeType
	//CylinderType
	//SphereType
)

// physic
const (
	MassType = iota
	ForceGravityType
	ImpulseType
	//KineticEnergyType
	//PotentialEnergyType
	//ForceWorkType
)

// mechanic
const (
	CompressiveStressType = iota
	RelativeDeformationType
	YoungsModulusType
)

// thermal
const (
	AmountHeatType = iota
	ThermalExpansionType
)

// other
const (
	gConst = 9.8
	//piConst = math.Pi
)
