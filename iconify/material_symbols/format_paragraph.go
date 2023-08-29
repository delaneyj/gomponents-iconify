package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatParagraph(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 20v-6q-2.075 0-3.538-1.463T4 9q0-2.075 1.463-3.538T9 4h9v2h-2v14h-2V6h-3v14H9Z"/>`),
		g.Group(children),
	)
}