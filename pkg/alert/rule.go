package alert

type Rule interface {
	CreateAlert() bool
	GetErrorName() string
	GetErrorData() interface{}
	GetReviewer() string
}

type MaxUsageRule struct {
	MaxUsage  float32
	ErrorName string
	GetData   func() float32
	SendTo    string
	data      float32
}

func (max *MaxUsageRule) CreateAlert() bool {
	data := max.GetData()
	if data >= max.MaxUsage {
		max.data = data
		return true
	}

	return false
}

func (max *MaxUsageRule) GetErrorName() string {
	return max.ErrorName
}

func (max *MaxUsageRule) GetErrorData() interface{} {
	return max.data
}

func (max *MaxUsageRule) GetReviewer() string {
	return max.SendTo
}
