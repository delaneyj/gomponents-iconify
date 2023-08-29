package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PianoSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21V3h18v18H3Zm2-2h3.25v-4.5H7V5H5v14Zm10.75 0H19V5h-2v9.5h-1.25V19Zm-6 0h4.5v-4.5H13V5h-2v9.5H9.75V19Z"/>`),
		g.Group(children),
	)
}