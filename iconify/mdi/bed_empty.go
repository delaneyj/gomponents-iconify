package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BedEmpty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 7H5v7H3V5H1v15h2v-3h18v3h2v-9a4 4 0 0 0-4-4"/>`),
		g.Group(children),
	)
}