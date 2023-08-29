package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrowseActivityOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 9V3h20v6h-2V5H4v4H2Zm0 9v-7h2v5h16v-5h2v7H2Zm0-7V9h6.6l1.475 2.875L13.425 6h1.2l1.5 3H22v2h-7.125l-.925-1.875L10.575 15h-1.2l-2-4H2ZM1 21v-2h22v2H1Zm11-10.5Z"/>`),
		g.Group(children),
	)
}