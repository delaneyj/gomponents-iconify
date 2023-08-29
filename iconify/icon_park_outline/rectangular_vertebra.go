package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RectangularVertebra(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m12 9l12-5l12 5l8 29l-20 6l-20-6l8-29Zm12 35V14m12-5l-12 5M12 9l12 5M5 38l19-6m20 6l-20-6"/>`),
		g.Group(children),
	)
}