package prey

import "testdoubles/positioner"

func NewPositionerStub() *PreyStub {
	return &PreyStub{}
}

type PreyStub struct {
	GetSpeedFunc    func() float64
	GetPositionFunc func() *positioner.Position
}

func (p *PreyStub) GetSpeed() float64 {
	return p.GetSpeedFunc()
}

func (p *PreyStub) GetPosition() *positioner.Position {
	return p.GetPositionFunc()
}
