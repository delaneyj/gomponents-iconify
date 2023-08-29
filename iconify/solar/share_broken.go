package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShareBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M4 12a2.5 2.5 0 1 0 2.5-2.5m7.5-3L9 10m5 7.5L9 14m7.5 7a2.5 2.5 0 1 0-2.5-2.5m4.665-11.75a2.5 2.5 0 1 1-.915-3.415"/>`),
		g.Group(children),
	)
}