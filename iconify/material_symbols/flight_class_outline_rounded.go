package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlightClassOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 13q-.825 0-1.413-.588T12 11V6q0-.825.588-1.413T14 4h2q.825 0 1.413.588T18 6v5q0 .825-.588 1.413T16 13h-2Zm0-2h2V6h-2v5Zm-4.5 7q-.675 0-1.2-.388t-.725-1.037l-2.5-8.3Q5.025 8.15 5.012 8T5 7.7V5q0-.425.288-.713T6 4q.425 0 .713.288T7 5v3l2.5 8H17q.425 0 .713.288T18 17q0 .425-.288.713T17 18H9.5ZM9 21q-.425 0-.713-.288T8 20q0-.425.288-.713T9 19h8q.425 0 .713.288T18 20q0 .425-.288.713T17 21H9Zm5-15h2h-2Z"/>`),
		g.Group(children),
	)
}