package calculator

import "reflect"

type CalcService struct {
	SpecificCalc map[string]interface{}
	CurrentCalc  interface{}
}

func CreateCalcService() *CalcService {
	specificCalc := make(map[string]interface{})
	math := make(map[string]interface{})
	math["cube"] = Math{}

	specificCalc["math"] = Math{}
	specificCalc["phys"] = Phys{}
	specificCalc["mech"] = Mech{}
	specificCalc["therm"] = Therm{}

	return &CalcService{
		SpecificCalc: specificCalc,
	}
}

func (cs CalcService) GetCount() int {
	return len(cs.SpecificCalc)
}

func (cs CalcService) GetNames() []string {
	names := make([]string, len(cs.SpecificCalc))
	for k, _ := range cs.SpecificCalc {
		names = append(names, k)
	}

	return names
}

func (cs CalcService) GetNamesInterface() []string {
	t := reflect.TypeOf(&cs.CurrentCalc).Elem()
	names := make([]string, t.NumMethod())

	for i := range t.NumMethod() {
		names = append(names, t.Method(i).Name)
	}

	return names
}

func (cs CalcService) GetCountInterface() int {
	t := reflect.TypeOf(&cs.CurrentCalc).Elem()

	return t.NumMethod()
}
