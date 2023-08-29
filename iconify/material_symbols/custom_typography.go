package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CustomTypography(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 22v-6h2v2h8v2h-8v2h-2Zm-8-2v-2h6v2H3Zm3.425-6H8.5l1.1-3.075h4.825L15.5 14h2.075l-4.5-12h-2.15l-4.5 12ZM10.2 9.2l1.75-4.975h.1L13.8 9.2h-3.6Z"/>`),
		g.Group(children),
	)
}