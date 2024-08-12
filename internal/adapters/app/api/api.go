package api

import "github.com/mnkrana/hex/internal/ports"

type Adapter struct {
	db    ports.DbPort
	arith ports.ArithmaticPort
}

func NewAdaptor(db ports.DbPort, arith ports.ArithmaticPort) *Adapter {
	return &Adapter{db: db, arith: arith}
}

func (apiAdapter Adapter) GetAddition(a int32, b int32) (int32, error) {
	answer, err := apiAdapter.arith.Addition(a, b)
	if err != nil {
		return 0, err
	}
	err = apiAdapter.db.AddToHistory(answer, "Add")
	if err != nil {
		return 0, err
	}
	return answer, nil
}

func (apiAdapter Adapter) GetSubtraction(a int32, b int32) (int32, error) {
	answer, err := apiAdapter.arith.Subtraction(a, b)
	if err != nil {
		return 0, err
	}
	err = apiAdapter.db.AddToHistory(answer, "Sub")
	if err != nil {
		return 0, err
	}
	return answer, nil
}

func (apiAdapter Adapter) GetMultiplication(a int32, b int32) (int32, error) {
	answer, err := apiAdapter.arith.Multiplication(a, b)
	if err != nil {
		return 0, err
	}
	err = apiAdapter.db.AddToHistory(answer, "Mul")
	if err != nil {
		return 0, err
	}
	return answer, nil
}

func (apiAdapter Adapter) GetDivision(a int32, b int32) (int32, error) {
	answer, err := apiAdapter.arith.Division(a, b)
	if err != nil {
		return 0, err
	}
	err = apiAdapter.db.AddToHistory(answer, "Div")
	if err != nil {
		return 0, err
	}
	return answer, nil
}
