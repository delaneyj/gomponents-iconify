package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Card(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 3h12v2h4v14h-4v2H6v-2H2V5h4V3Zm0 4H4v10h2V7Zm2 12h8V5H8v14ZM18 7v10h2V7h-2Z"/>`),
		g.Group(children),
	)
}