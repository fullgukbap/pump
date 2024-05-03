// exr 패키지는 exchnage rate의 약자로 환율 정보와 관련된 모든것을 구현합니다.
package exr

import (
	"time"
)

// New 함수는 Exr 구조체의 생성자 함수 입니다.
func New(adapter exchnageRateAble) *exr {
	return &exr{
		exchnageRateAble: adapter,
		points:           make([]point, 0),
	}
}

// exr 구조체는 모든 환율 정보를 포함하고 있습니다.
type exr struct {
	exchnageRateAble
	points []point
}

// point 구조체는 환율 정보의 가장 작은 단위 입니다.
type point struct {
	CurCode      string
	DealBaseRate float64
	CreatedAt    time.Time
}
