package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AnchorSimpleBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M224 108h-24a12 12 0 0 0 0 24h11.15A84.21 84.21 0 0 1 140 203.14V97.94a36 36 0 1 0-24 0v105.2A84.21 84.21 0 0 1 44.85 132H56a12 12 0 0 0 0-24H32a12 12 0 0 0-12 12a108 108 0 0 0 216 0a12 12 0 0 0-12-12Zm-96-56a12 12 0 1 1-12 12a12 12 0 0 1 12-12Z"/>`),
		g.Group(children),
	)
}