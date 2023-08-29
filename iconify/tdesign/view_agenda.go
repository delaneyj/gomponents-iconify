package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewAgenda(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 3h20v8H2V3Zm2 2v4h16V5H4Zm-2 8h20v8H2v-8Zm2 2v4h16v-4H4Z"/>`),
		g.Group(children),
	)
}