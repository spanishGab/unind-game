package utils

func AreAllPositive[V float32 | float64 | int | uint](value ...V) bool {
	for _, attributePoint := range value {
		if attributePoint < 0 {
			return false
		}
	}
	return true
}
