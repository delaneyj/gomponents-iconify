package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TagSimpleBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m250 121.34l-45.64-68.43A20 20 0 0 0 187.72 44H40a20 20 0 0 0-20 20v128a20 20 0 0 0 20 20h147.72a20 20 0 0 0 16.64-8.91L250 134.66a12 12 0 0 0 0-13.32ZM185.58 188H44V68h141.58l40 60Z"/>`),
		g.Group(children),
	)
}