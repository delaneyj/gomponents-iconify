package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VerticalShadesOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 21v-2h2V3h16v16h2v2H2Zm4-2h2V5H6v14Zm4 0h4V5h-4v14Zm6 0h2V5h-2v14ZM6 19V5v14Zm12 0V5v14Z"/>`),
		g.Group(children),
	)
}