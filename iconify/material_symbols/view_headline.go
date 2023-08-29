package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewHeadline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 15v-2h16v2H4Zm0 4v-2h16v2H4Zm0-8V9h16v2H4Zm0-4V5h16v2H4Z"/>`),
		g.Group(children),
	)
}