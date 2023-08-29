package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PrintDisabled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.775 22.625L18 20.85V21H6v-4H2v-6q0-1.275.875-2.138T5 8h.15L1.375 4.225L2.8 2.8l18.4 18.4l-1.425 1.425ZM8 19h8v-.15L12.15 15H8v4Zm11.85-2l-9-9H19q1.275 0 2.138.863T22 11v6h-2.15ZM16 7H9.85l-4-4H18v4h-2Zm2 5.5q.425 0 .713-.288T19 11.5q0-.425-.288-.713T18 10.5q-.425 0-.713.288T17 11.5q0 .425.288.713T18 12.5Z"/>`),
		g.Group(children),
	)
}