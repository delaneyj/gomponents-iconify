package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowRotaryFirstLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 10a3 3 0 1 1 0-6a3 3 0 0 1 0 6zm0 0v10M13.5 9.5L5 18m5 0H5v-5"/>`),
		g.Group(children),
	)
}