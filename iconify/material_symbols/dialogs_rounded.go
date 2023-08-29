package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DialogsRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 16h6q.425 0 .713-.288T16 15V9q0-.425-.288-.713T15 8H9q-.425 0-.713.288T8 9v6q0 .425.288.713T9 16Zm-4 5q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v14q0 .825-.588 1.413T19 21H5Z"/>`),
		g.Group(children),
	)
}