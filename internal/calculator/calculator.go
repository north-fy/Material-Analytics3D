package calculator

type CalcService struct {
	specificCalc map[string]interface{}
}

func CreateCalcService() *CalcService {
	specificCalc := make(map[string]interface{})
	specificCalc["math"] = Math{}
	specificCalc["phys"] = Phys{}
	specificCalc["mech"] = Mech{}
	specificCalc["therm"] = Therm{}

	return &CalcService{
		specificCalc: specificCalc,
	}
}
