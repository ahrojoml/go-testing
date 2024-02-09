package simulator

func NewCatchSimulatorMock() *CatchSimulatorMock {
	return &CatchSimulatorMock{}
}

type CatchSimulatorMock struct {
	CanCatchFunc func(hunter, prey *Subject) bool
	Calls        struct {
		CanCatch int
	}
}

func (cs *CatchSimulatorMock) CanCatch(hunter, prey *Subject) bool {
	cs.Calls.CanCatch++
	return cs.CanCatchFunc(hunter, prey)
}
