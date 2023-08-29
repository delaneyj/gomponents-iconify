package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PanTool(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.05 23q-.75 0-1.4-.338T7.575 21.7L1.2 12.375l.6-.575q.475-.475 1.125-.55t1.175.3L7 13.575V4q0-.425.288-.713T8 3q.425 0 .713.288T9 4v8h2V2q0-.425.288-.713T12 1q.425 0 .713.288T13 2v10h2V3q0-.425.288-.713T16 2q.425 0 .713.288T17 3v9h2V5q0-.425.288-.713T20 4q.425 0 .713.288T21 5v14q0 1.65-1.175 2.825T17 23h-6.95Z"/>`),
		g.Group(children),
	)
}