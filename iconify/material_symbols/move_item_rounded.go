package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoveItemRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.15 13H9q-.425 0-.713-.288T8 12q0-.425.288-.713T9 11h11.15l-.85-.85q-.3-.3-.288-.7t.288-.7q.3-.3.712-.312t.713.287L23.3 11.3q.3.3.3.7t-.3.7l-2.575 2.575q-.3.3-.713.288t-.712-.313q-.275-.3-.288-.7t.288-.7l.85-.85ZM15 8V5H5v14h10v-3q0-.425.288-.713T16 15q.425 0 .713.288T17 16v3q0 .825-.588 1.413T15 21H5q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h10q.825 0 1.413.588T17 5v3q0 .425-.288.713T16 9q-.425 0-.713-.288T15 8Z"/>`),
		g.Group(children),
	)
}