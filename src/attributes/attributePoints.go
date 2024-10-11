package attributes

import "errors"

type Attribute string

const (
	ATTACK Attribute = "attack"
	DEFENSE Attribute = "defense"
	HEALTH Attribute = "healh"
	INTELLIGENCE Attribute = "intelligence"
)

type AttributePoints struct {
	points map[Attribute]float64
}

func (ap *AttributePoints) Get(attribute Attribute) (float64, error) {
	value, ok := ap.points[attribute]
	if !ok {
		return 0.0, errors.New("attribute not found")
	}
	return value, nil
}

func (ap *AttributePoints) Has(attribute Attribute) bool {
	_, ok := ap.points[attribute]
	return ok
}