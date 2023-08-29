package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PortUsbCTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M16 10.5a1.5 1.5 0 1 1 0 3H8a1.5 1.5 0 1 1 0-3h8ZM16 9H8a3 3 0 1 0 0 6h8a3 3 0 0 0 0-6Z"/>`),
		g.Group(children),
	)
}