package FilterType

import (
	filter "github.com/Antipascal/image-generator/pkg/img/filters"
	"image/color"
)

const (
	None     = "🌈 None"
	Red      = "🔴 Red"
	Grey     = "⬜️ Grey"
	Purple   = "🟣 Purple"
	DarkBlue = "🔵 Dark blue"
	Yellow   = "🟡 Yellow"
	Orange   = "🟠 Orange"
	Green    = "🟢 Green"
	White    = "⚪️ White"
	Black    = "⚫️ Black"
)

func GetFilterMethod(filterName string) func(color.RGBA) color.RGBA {
	switch filterName {
	case Red:
		return filter.RedFilter
	case Grey:
		return filter.GreyFilter
	case Purple:
		return filter.PurpleFilter
	case DarkBlue:
		return filter.DarkBlueFilter
	case Yellow:
		return filter.YellowFilter
	case Orange:
		return filter.OrangeFilter
	case Green:
		return filter.GreenFilter
	case White:
		return filter.WhiteFilter
	case Black:
		return filter.BlackFilter
	}
	return nil
}
