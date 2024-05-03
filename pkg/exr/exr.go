// exr 패키지는 exchnage rate의 약자로 환율 정보를 구현합니다.
package exr

// New 함수는 Exr 구조체의 생성자 함수 입니다.
func New(adapter exchnageRateAble) *exr {
	return &exr{
		exchnageRateAble: adapter,
	}
}

// Exr 구조체는 모든 환율 정보를 포함하고 있습니다.
type exr struct {
	exchnageRateAble
}
