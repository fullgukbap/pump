package exr

// ExchnageRateAble 인터페이스는 환율 정보를 얻는 행위들을 정의합니다.
type exchnageRateAble interface {
	// GetDealBaseRate 함수는 통화코드(currencyCode)에 해당되는 매매 기준율을 반환홥니다.
	// 2024.05.03 기준으로 지원한 통화코드는 미국, 일본, 호주, 캐나다 입니다.
	GetDealBaseRate(currencyCode string) (float64, error)
}
