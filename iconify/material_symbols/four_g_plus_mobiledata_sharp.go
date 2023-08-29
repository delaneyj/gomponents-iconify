package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FourGPlusMobiledataSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 15v-2h-2v-2h2V9h2v2h2v2h-2v2h-2ZM5 17v-3H1V7h2v5h2V7h2v5h1v2H7v3H5Zm4 0V7h8v2h-6v6h4v-2h-2v-2h4v6H9Z"/>`),
		g.Group(children),
	)
}