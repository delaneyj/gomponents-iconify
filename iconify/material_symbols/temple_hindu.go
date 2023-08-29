package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TempleHindu(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m7.8 7l1.175-3.875V1h2v2H13V1h2v2l1.2 4H7.8ZM2 22V11h2v2h16v-2h2v11h-9v-5h-2v5H2Zm4.6-11l.6-2h9.6l.6 2H6.6Z"/>`),
		g.Group(children),
	)
}