package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignJustifySpaceBetween(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 22v-5h-3V7h3V2h2v20h-2ZM2 22V2h2v5h3v10H4v5H2Z"/>`),
		g.Group(children),
	)
}