package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatInkHighlighterOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 24v-4h20v4H2Zm12.6-12l-1.3-1.3L12 9.4l-4 4l2.575 2.6l4.025-4Zm-1.175-4l1.3 1.275l1.275 1.3L20 6.6L17.4 4l-3.975 4ZM3.5 18l3.15-3.15l-.75-.725V12.7l5.425-5.425l5.4 5.4L11.3 18.1H9.9l-.75-.75l-.65.65h-5Zm7.825-10.725l6.1-6.1L22.85 6.6l-6.125 6.075l-5.4-5.4Z"/>`),
		g.Group(children),
	)
}