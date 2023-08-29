package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MagnifyDockedSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 20V4h20v16H2Zm2-4h16V6H4v10Zm10-2h2v-2h2v-2h-2V8h-2v2h-2v2h2v2Z"/>`),
		g.Group(children),
	)
}