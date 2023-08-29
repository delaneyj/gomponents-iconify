package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeartSmallOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="m6.5 5.5l1 1l1-1a1.414 1.414 0 1 1 2 2l-3 3l-3-3a1.414 1.414 0 1 1 2-2Z"/>`),
		g.Group(children),
	)
}