package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TickSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M14.707 3L5.5 12.207L.293 7L1 6.293l4.5 4.5l8.5-8.5l.707.707Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}