package utils

const (
	USD = "USD"
	GPD = "GPD"
	EUR = "EUR"
	CAD = "CAD"
)

func IsSurportedCurreny(currency string) bool {
	switch currency {
	case USD, CAD, EUR, GPD:
		return true
	}
	return false
}
