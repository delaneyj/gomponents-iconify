package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextRotationLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m13 29l3.667-8M35 29l-3.667-8m0 0L29.5 17L24 5l-5.5 12l-1.833 4m14.666 0H16.667M42 37H6l6 6"/>`),
		g.Group(children),
	)
}