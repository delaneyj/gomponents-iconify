package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Luggage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 21q-.825 0-1.413-.588T5 19V8q0-.825.588-1.413T7 6h2V3.5q0-.625.438-1.063T10.5 2h3q.625 0 1.063.438T15 3.5V6h2q.825 0 1.413.588T19 8v11q0 .825-.588 1.413T17 21q0 .425-.288.713T16 22q-.425 0-.713-.288T15 21H9q0 .425-.288.713T8 22q-.425 0-.713-.288T7 21Zm1-3h1.5V9H8v9Zm3.25 0h1.5V9h-1.5v9Zm3.25 0H16V9h-1.5v9Zm-4-12h3V3.5h-3V6Z"/>`),
		g.Group(children),
	)
}