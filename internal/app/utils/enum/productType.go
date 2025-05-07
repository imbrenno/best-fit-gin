package enum

import "fmt"

type Type string

const (
	NON_ALCOHOLIC Type = "NON_ALCOHOLIC"
	ALCOHOLIC     Type = "ALCOHOLIC"
)

func ValidadeProductType(productType string) error {
	switch productType {
	case string(NON_ALCOHOLIC), string(ALCOHOLIC):
		return nil
	default:
		return fmt.Errorf("invalid product type: %s", productType)
	}
}
