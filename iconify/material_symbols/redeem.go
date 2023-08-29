package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Redeem(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 17v2h16v-2H4ZM4 6h2.2q-.125-.225-.163-.475T6 5q0-1.25.875-2.125T9 2q.75 0 1.388.388t1.112.962L12 4l.5-.65q.45-.6 1.1-.975T15 2q1.25 0 2.125.875T18 5q0 .275-.038.525T17.8 6H20q.825 0 1.412.588T22 8v11q0 .825-.588 1.413T20 21H4q-.825 0-1.413-.588T2 19V8q0-.825.588-1.413T4 6Zm0 8h16V8h-5.1l2.1 2.85L15.4 12L12 7.4L8.6 12L7 10.85L9.05 8H4v6Zm5-8q.425 0 .713-.288T10 5q0-.425-.288-.713T9 4q-.425 0-.713.288T8 5q0 .425.288.713T9 6Zm6 0q.425 0 .713-.288T16 5q0-.425-.288-.713T15 4q-.425 0-.713.288T14 5q0 .425.288.713T15 6Z"/>`),
		g.Group(children),
	)
}