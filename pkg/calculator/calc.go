package calculator

// Calculator 提供基础的计算功能
type Calculator struct {
	precision int
}

// NewCalculator 创建一个新的计算器实例
func NewCalculator(precision int) *Calculator {
	return &Calculator{
		precision: precision,
	}
}

// Add 执行加法运算
func (c *Calculator) Add(a, b float64) float64 {
	return a + b
}

// Multiply 执行乘法运算
func (c *Calculator) Multiply(a, b float64) float64 {
	return a * b
}
