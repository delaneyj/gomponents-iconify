package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ComponentInput(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 7h20v9H2V7Zm2 2v5h16V9H4Zm4 1v3H6v-3h2Z"/>`),
		g.Group(children),
	)
}