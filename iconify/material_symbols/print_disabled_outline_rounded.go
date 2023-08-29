package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PrintDisabledOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 17H3q-.425 0-.713-.288T2 16v-5q0-1.275.875-2.138T5 8h.15L2.075 4.925q-.3-.3-.3-.7T2.1 3.5q.275-.275.7-.275t.725.275l16.975 17q.3.3.3.7t-.3.7q-.3.3-.713.3t-.712-.3l-2.5-2.5L18 18v1q0 .825-.588 1.413T16 21H8q-.825 0-1.413-.588T6 19v-2Zm2-2v4h8v-.15L12.15 15H8Zm11.85 2l-2-2H20v-4q0-.425-.288-.713T19 10h-6.15l-2-2H19q1.275 0 2.138.863T22 11v5q0 .425-.288.713T21 17h-1.15ZM16 8V5H7.85l-2-2H17q.425 0 .713.288T18 4v4h-2ZM4 15h2v-2h4.15l-3-3H5q-.425 0-.713.288T4 11v4Zm14-2.5q.425 0 .713-.288T19 11.5q0-.425-.288-.713T18 10.5q-.425 0-.713.288T17 11.5q0 .425.288.713T18 12.5ZM5 10h2.15H4h1Zm14 0h1h-7.15H19Z"/>`),
		g.Group(children),
	)
}