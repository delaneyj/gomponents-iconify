package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PortUsbCTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M19 12a2.996 2.996 0 0 1-3 3H8a3 3 0 0 1 0-6h8a2.996 2.996 0 0 1 3 3Z"/>`),
		g.Group(children),
	)
}