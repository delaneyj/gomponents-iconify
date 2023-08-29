package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PowerSetting(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M64 472v-43h43v43H64zm85 0v-43h43v43h-43zM192 3v213h-43V3h43zm76 52q34 23 53.5 60t19.5 80q0 70-50 120t-120.5 50T50 315T0 195q0-43 19.5-80T73 55l31 30q-28 18-44.5 47T43 195q0 53 37.5 90.5T171 323t90.5-37.5T299 195q0-34-17-63t-45-46zm-33 417v-43h42v43h-42z"/>`),
		g.Group(children),
	)
}