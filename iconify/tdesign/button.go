package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Button(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 5h20v14H2V5Zm2 2v10h16V7H4Zm2 2h12v6H6V9Zm2 2v2h8v-2H8Z"/>`),
		g.Group(children),
	)
}