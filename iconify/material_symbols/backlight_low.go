package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BacklightLow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 15v-2h3v2H2Zm4.35-5.25L4.225 7.625l1.4-1.4L7.75 8.35l-1.4 1.4ZM7 18v-3h10v3H7Zm4-11V4h2v3h-2Zm6.65 2.775l-1.4-1.425l2.125-2.125l1.4 1.425l-2.125 2.125ZM19 15v-2h3v2h-3Z"/>`),
		g.Group(children),
	)
}