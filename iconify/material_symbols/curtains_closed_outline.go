package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurtainsClosedOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 21v-2h2V3h16v16h2v2H2Zm4-2h3V5H6v14Zm5 0h2V5h-2v14Zm4 0h3V5h-3v14Zm-9 0V5v14Zm12 0V5v14Z"/>`),
		g.Group(children),
	)
}