package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InformationSlabSymbol(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 9h-2V7h2v2m1 6v2h-4v-2h1v-2h-1v-2h3v4h1Z"/>`),
		g.Group(children),
	)
}