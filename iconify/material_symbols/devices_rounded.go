package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DevicesRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3.5 20q-.625 0-1.063-.438T2 18.5q0-.625.438-1.063T3.5 17H4V6q0-.825.588-1.413T6 4h14q.425 0 .713.288T21 5q0 .425-.288.713T20 6H6v11h4.5q.625 0 1.063.438T12 18.5q0 .625-.438 1.063T10.5 20h-7ZM15 20q-.425 0-.713-.288T14 19V9q0-.425.288-.713T15 8h6q.425 0 .713.288T22 9v10q0 .425-.288.713T21 20h-6Z"/>`),
		g.Group(children),
	)
}