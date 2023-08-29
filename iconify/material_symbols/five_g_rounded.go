package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FiveGRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 17H4q-.425 0-.713-.288T3 16q0-.425.288-.713T4 15h4v-2H4q-.425 0-.713-.288T3 12V8q0-.425.288-.713T4 7h5q.425 0 .713.288T10 8q0 .425-.288.713T9 9H5v2h3q.825 0 1.413.588T10 13v2q0 .825-.588 1.413T8 17Zm6 0q-.825 0-1.413-.588T12 15V9q0-.825.588-1.413T14 7h6q.425 0 .713.288T21 8q0 .425-.288.713T20 9h-6v6h5v-2h-1.5q-.425 0-.713-.288T16.5 12q0-.425.288-.713T17.5 11H20q.425 0 .713.288T21 12v3q0 .825-.588 1.413T19 17h-5Z"/>`),
		g.Group(children),
	)
}