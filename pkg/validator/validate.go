package validator

// Validator 提供数据验证功能
type Validator struct {
	maxValue float64
}

// NewValidator 创建一个新的验证器实例
func NewValidator(maxValue float64) *Validator {
	return &Validator{
		maxValue: maxValue,
	}
}

// IsValid 检查值是否有效
func (v *Validator) IsValid(value float64) bool {
	return value <= v.maxValue
}
