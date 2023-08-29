package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatInkHighlighterSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 24v-4h20v4H2Zm1.5-6l3.15-3.15l-.75-.725V12.7L10.6 8l5.4 5.425l-4.7 4.675H9.9l-.75-.75l-.65.65h-5ZM12 6.575l5.425-5.4l5.4 5.425l-5.4 5.4L12 6.575Z"/>`),
		g.Group(children),
	)
}