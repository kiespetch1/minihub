package metrics

type Metrics interface {
	Base() *BaseMetrics
}

type DefaultMetrics struct {
	*BaseMetrics
}

func NewDefault(namespace, serviceName string) *DefaultMetrics {
	return &DefaultMetrics{
		BaseMetrics: NewBaseMetrics(namespace, serviceName),
	}
}

func (m *DefaultMetrics) Base() *BaseMetrics {
	return m.BaseMetrics
}
