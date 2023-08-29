package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowRotaryLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 10a3 3 0 1 1 0-6a3 3 0 0 1 0 6zm0 0v10M13 7H3m4 4L3 7l4-4"/>`),
		g.Group(children),
	)
}