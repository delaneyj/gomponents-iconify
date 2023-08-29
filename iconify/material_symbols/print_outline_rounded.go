package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PrintOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 8V5H8v3H6V4q0-.425.288-.713T7 3h10q.425 0 .713.288T18 4v4h-2ZM4 10h16H4Zm14 2.5q.425 0 .713-.288T19 11.5q0-.425-.288-.713T18 10.5q-.425 0-.713.288T17 11.5q0 .425.288.713T18 12.5ZM8 19h8v-4H8v4Zm0 2q-.825 0-1.413-.588T6 19v-2H3q-.425 0-.713-.288T2 16v-5q0-1.275.875-2.138T5 8h14q1.275 0 2.138.863T22 11v5q0 .425-.288.713T21 17h-3v2q0 .825-.588 1.413T16 21H8Zm12-6v-4q0-.425-.288-.713T19 10H5q-.425 0-.713.288T4 11v4h2v-2h12v2h2Z"/>`),
		g.Group(children),
	)
}