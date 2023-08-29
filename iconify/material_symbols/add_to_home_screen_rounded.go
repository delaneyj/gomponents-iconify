package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AddToHomeScreenRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 23q-.825 0-1.413-.588T6 21v-4h2v1h10V6H8v1H6V3q0-.825.588-1.413T8 1h10q.825 0 1.413.588T20 3v18q0 .825-.588 1.413T18 23H8Zm2-11.6l-4.9 4.9q-.275.275-.7.275t-.7-.275q-.275-.275-.275-.7t.275-.7L8.6 10H6q-.425 0-.713-.288T5 9q0-.425.288-.713T6 8h5q.425 0 .713.288T12 9v5q0 .425-.288.713T11 15q-.425 0-.713-.288T10 14v-2.6Z"/>`),
		g.Group(children),
	)
}