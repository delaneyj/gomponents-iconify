package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpHourglassTop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m6 2l.01 6L10 12l-3.99 4.01L6 22h12v-6l-4-4l4-3.99V2H6zm10 14.5V20H8v-3.5l4-4l4 4z"/>`),
		g.Group(children),
	)
}