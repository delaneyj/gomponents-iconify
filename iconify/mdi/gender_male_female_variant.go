package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GenderMaleFemaleVariant(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 3a5 5 0 0 0 2 4a5 5 0 0 0-2 4a5 5 0 0 0 4 4.9V18H9v2h2v2h2v-2h2v-2h-2v-2.1a5 5 0 0 0 4-4.9a5 5 0 0 0-2-4a5 5 0 0 0 2-4h-2a3 3 0 0 1-3 3a3 3 0 0 1-3-3m3 5a3 3 0 0 1 3 3a3 3 0 0 1-3 3a3 3 0 0 1-3-3a3 3 0 0 1 3-3Z"/>`),
		g.Group(children),
	)
}