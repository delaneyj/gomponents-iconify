package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tune(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 21v-6h2v2h8v2h-8v2h-2Zm-8-2v-2h6v2H3Zm4-4v-2H3v-2h4V9h2v6H7Zm4-2v-2h10v2H11Zm4-4V3h2v2h4v2h-4v2h-2ZM3 7V5h10v2H3Z"/>`),
		g.Group(children),
	)
}