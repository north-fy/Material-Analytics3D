package calculator

// масса тела m = qV
// вес тела/сила тяжести P = mg (обязательно переводить в метры)
// импульс тела P = mv
// Кин. энергия E_k = (mv^2) / 2
// Пот. энергия E_п = mgh
// Работа силы A = FScosa

type PhysicProperties struct {
	Mass         float64
	ForceGravity float64
	Impulse      float64
	//KineticEnergy   float64
	//PotentialEnergy float64
	//ForceWork       float64
}

type PhysicCalculator interface {
	Calculate() *PhysicProperties
	GetType() int
}

// Mass расчет массы тела по формуле m = qV
// Определенный тип - 0
type Mass struct {
	Density float64
	Volume  float64
}

func NewMass(q, v float64) *Mass {
	return &Mass{Density: q, Volume: v}
}

func (m *Mass) Calculate() *PhysicProperties {
	mass := m.Density * m.Volume
	return &PhysicProperties{Mass: mass}
}

func (m *Mass) GetType() int {
	return MassType
}

// ForceGravity расчет силы тяжести/вес тела по формуле F = mg (P = mg)
// Определенный тип - 1
type ForceGravity struct {
	Mass float64
}

func NewForceGravity(m float64) *ForceGravity {
	return &ForceGravity{Mass: m}
}

func (f *ForceGravity) Calculate() *PhysicProperties {
	forceGravity := f.Mass * gConst
	return &PhysicProperties{ForceGravity: forceGravity}
}

func (f *ForceGravity) GetType() int {
	return ForceGravityType
}

// Impulse расчет импульса тела по формуле P = mv
// Определенный тип - 2
type Impulse struct {
	Mass     float64
	Velocity float64
}

func NewImpulse(m, v float64) *Impulse {
	return &Impulse{Mass: m, Velocity: v}
}

func (i *Impulse) Calculate() *PhysicProperties {
	impulse := i.Mass * i.Velocity
	return &PhysicProperties{Impulse: impulse}
}

func (i *Impulse) GetType() int {
	return ImpulseType
}

func CreatePhysicCalculator(physicType int, values ...float64) (PhysicCalculator, error) {
	switch physicType {
	case 0: // Mass m = qV
		if len(values) == 2 {
			return NewMass(values[0], values[1]), nil
		}
		return nil, ErrEnoughArg

	case 1: // ForceGravity F = mg
		if len(values) == 1 {
			return NewForceGravity(values[0]), nil
		}
		return nil, ErrEnoughArg

	case 2: // Impulse I = mv
		if len(values) == 2 {
			return NewImpulse(values[0], values[1]), nil
		}
		return nil, ErrEnoughArg

	default: // Если тип неизвестен
		return nil, ErrUnknownType
	}
}
