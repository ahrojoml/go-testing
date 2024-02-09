package positioner

func NewPositionerStub() *PositionerStub {
	return &PositionerStub{}
}

type PositionerStub struct {
	GetLinearDistanceFunc func(from, to *Position) float64
}

func (p *PositionerStub) GetLinearDistance(from, to *Position) float64 {
	return p.GetLinearDistanceFunc(from, to)
}
