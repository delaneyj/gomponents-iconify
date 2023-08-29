package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrowseActivitySharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 21v-2h22v2H1Zm1-3v-7h5.375l2 4h1.2l3.375-5.875l.925 1.875H22v7H2Zm8.075-6.125L8.6 9H2V3h20v6h-5.875l-1.5-3h-1.2l-3.35 5.875Z"/>`),
		g.Group(children),
	)
}