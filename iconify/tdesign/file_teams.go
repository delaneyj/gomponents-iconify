package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileTeams(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 1h12.414L21 6.586V23H3V1Zm2 2v18h14V9h-6V3H5Zm10 .414V7h3.586L15 3.414ZM9 10h6v2h-2v6h-2v-6H9v-2Z"/>`),
		g.Group(children),
	)
}