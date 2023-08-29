package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OpenInNewDownRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v6q0 .425-.288.713T20 12q-.425 0-.713-.288T19 11V5H5v14h6q.425 0 .713.288T12 20q0 .425-.288.713T11 21H5Zm12.6-2L9 10.4q-.275-.275-.288-.688T9 9q.275-.275.7-.275t.7.275l8.6 8.575V15q0-.425.288-.713T20 14q.425 0 .713.288T21 15v5q0 .425-.288.713T20 21h-5q-.425 0-.713-.288T14 20q0-.425.288-.713T15 19h2.6Z"/>`),
		g.Group(children),
	)
}