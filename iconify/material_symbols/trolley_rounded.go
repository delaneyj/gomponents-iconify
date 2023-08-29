package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrolleyRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 17H6q-.825 0-1.413-.588T4 15V5H3q-.425 0-.713-.288T2 4q0-.425.288-.713T3 3h1q.825 0 1.413.588T6 5v10h14q.425 0 .713.288T21 16q0 .425-.288.713T20 17ZM6 22q-.825 0-1.413-.588T4 20q0-.825.588-1.413T6 18q.825 0 1.413.588T8 20q0 .825-.588 1.413T6 22Zm2-8q-.425 0-.713-.288T7 13V9q0-.425.288-.713T8 8h4q.425 0 .713.288T13 9v4q0 .425-.288.713T12 14H8Zm7 0q-.425 0-.713-.288T14 13V9q0-.425.288-.713T15 8h4q.425 0 .713.288T20 9v4q0 .425-.288.713T19 14h-4Zm4 8q-.825 0-1.413-.588T17 20q0-.825.588-1.413T19 18q.825 0 1.413.588T21 20q0 .825-.588 1.413T19 22Z"/>`),
		g.Group(children),
	)
}