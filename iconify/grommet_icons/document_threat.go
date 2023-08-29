package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentThreat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M4.998 7V1H19.5L23 4.5V23h-6m1-22v5h5M9 23a5 5 0 1 0 0-10a5 5 0 0 0 0 10Zm0-12V9c0-1 0-2 2-2s2 1 2 2s0 2 2 2h2m-9 0h2v2H8v-2Z"/>`),
		g.Group(children),
	)
}