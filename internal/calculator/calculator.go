package calculator

import (
	"reflect"
)

/*
Здесь оставлю свое повествование
Короче похуй, надо либо переделывать архитектуру калькулятора
Либо забивать болт и создавать каждый кальк отдельно, а потом значения прикручивать
Мол у нас есть дата map[string]interface{calculate, gettype, set_value}
Через дату узнавать какой кальк нужен, по интерфейсу обращаться к нему
gettype - узнать кол-во значений мб?
set_value очевидно, что установить values...float64 значений
calculate - дураку понятно
*/

type CalcService struct {
	SpecificCalc map[string]interface{}
	CurrentCalc  interface{}
}

func CreateCalcService() *CalcService {
	specificCalc := make(map[string]interface{})

	specificCalc["Math"] = Math{}
	specificCalc["Phys"] = Phys{}
	specificCalc["Mech"] = Mech{}
	specificCalc["Therm"] = Therm{}

	return &CalcService{
		SpecificCalc: specificCalc,
	}
}

func (cs CalcService) GetCount() int {
	return len(cs.SpecificCalc)
}

func (cs CalcService) GetNames() []string {
	names := []string{}
	for k, _ := range cs.SpecificCalc {
		names = append(names, k)
	}

	return names
}

func (cs *CalcService) SetCurrentCalc(key string) {
	switch cs.SpecificCalc[key].(type) {
	case Math:
		cs.CurrentCalc = Math{}
	case Phys:
		cs.CurrentCalc = Phys{}
	case Therm:
		cs.CurrentCalc = Therm{}
	case Mech:
		cs.CurrentCalc = Mech{}
	}
}

func (cs *CalcService) CalcCurrentCalc(typeCalc string, values map[string]float64) (map[string]float64, error) {
	answer := make(map[string]float64, 3)

	switch cs.CurrentCalc.(type) {
	case Math:
		mp, err := CreateMathCalculator(typeCalc, values)
		if err != nil {
			return nil, err
		}

		mpStat := mp.Calculate()
		// заменить потом
		answer["TotalArea"] = mpStat.TotalArea
		answer["Volume"] = mpStat.Volume

		return answer, nil

	case Phys:
		mp, err := CreatePhysicCalculator(typeCalc, values)
		if err != nil {
			return nil, err
		}

		mpStat := mp.Calculate()
		// заменить потом
		answer["Mass"] = mpStat.Mass
		answer["ForceGravity"] = mpStat.ForceGravity
		answer["Impulse"] = mpStat.Impulse

		return answer, nil

	case Mech:
		mp, err := CreateMechanicCalculator(typeCalc, values)
		if err != nil {
			return nil, err
		}

		mpStat := mp.Calculate()
		// заменить потом
		answer["CompressiveStress"] = mpStat.CompressiveStress
		answer["YoungsModulus"] = mpStat.YoungsModulus
		answer["RelativeDeformation"] = mpStat.RelativeDeformation

		return answer, nil

	case Therm:
		mp, err := CreateThermicCalculator(typeCalc, values)
		if err != nil {
			return nil, err
		}

		mpStat := mp.Calculate()
		// заменить потом
		answer["AmountHeat"] = mpStat.AmountHeat
		answer["ThermalExpansion"] = mpStat.ThermalExpansion

		return answer, nil
	}

	return nil, nil
}

func (cs CalcService) GetNamesInterface() []string {
	switch cs.CurrentCalc.(type) {
	case Math:
		return []string{"Cube", "Parallelepiped", "Pyramid"}
	case Phys:
		return []string{"Mass", "ForceGravity", "Impulse"}
	case Mech:
		return []string{"CompressiveStress", "RelativeDeformation", "YoungsModulus"}
	case Therm:
		return []string{"AmountHeat", "ThermalExpansion"}
	}
	return nil
}

func (cs CalcService) GetNamesInterfaceValues(method string) []string {
	values := make([]string, 0)

	switch method {
	// math
	case "Cube":
		values = append(values, "Side")
	case "Parallelepiped":
		values = append(values, "Lenght", "Width", "Height")
	case "Pyramid":
		values = append(values, "BaseSide", "Height")
	// phys
	case "Mass":
		values = append(values, "Density", "Volume")
	case "ForceGravity":
		values = append(values, "Mass")
	case "Impulse":
		values = append(values, "Mass", "Velocity")
	// mech
	case "CompressiveStress":
		values = append(values, "Force", "CrossArea")
	case "RelativeDeformation":
		values = append(values, "Lenght", "ChangeLenght")
	case "YoungsModulus":
		values = append(values, "RelativeDeformationValue", "CompressiveStressValue")
	// therm
	case "AmountHeat":
		values = append(values, "SpecificHeat", "Mass", "Temperature1", "Temperature2")
	case "ThermalExpansion":
		values = append(values, "Coefficient", "Lenght", "Temperature1", "Temperature2")
	}

	return values
}

func (cs CalcService) GetCountInterface() int {
	t := reflect.TypeOf(&cs.CurrentCalc).Elem()

	return t.NumMethod()
}
