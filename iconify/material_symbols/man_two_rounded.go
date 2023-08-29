package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ManTwoRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.5 21v-6H9q-.425 0-.713-.288T8 14V9q0-.825.588-1.413T10 7h4q.825 0 1.413.588T16 9v5q0 .425-.288.713T15 15h-1.5v6q0 .425-.288.713T12.5 22h-1q-.425 0-.713-.288T10.5 21ZM12 6q-.825 0-1.413-.588T10 4q0-.825.588-1.413T12 2q.825 0 1.413.588T14 4q0 .825-.588 1.413T12 6Z"/>`),
		g.Group(children),
	)
}