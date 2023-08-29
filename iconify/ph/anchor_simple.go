package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AnchorSimple(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M224 112h-24a8 8 0 0 0 0 16h15.64A88.15 88.15 0 0 1 136 207.63V95a32 32 0 1 0-16 0v112.63A88.15 88.15 0 0 1 40.36 128H56a8 8 0 0 0 0-16H32a8 8 0 0 0-8 8a104 104 0 0 0 208 0a8 8 0 0 0-8-8ZM112 64a16 16 0 1 1 16 16a16 16 0 0 1-16-16Z"/>`),
		g.Group(children),
	)
}