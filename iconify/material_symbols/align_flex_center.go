package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignFlexCenter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 22v-8.5H3v-3h8V2h2v8.5h8v3h-8V22h-2Z"/>`),
		g.Group(children),
	)
}