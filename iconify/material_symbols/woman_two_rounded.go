package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WomanTwoRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.5 22q-.425 0-.713-.288T10.5 21v-5H8.475q-.525 0-.825-.438t-.1-.937l2.5-6.325q.25-.6.775-.95T12 7q.65 0 1.175.35t.775.95l2.5 6.325q.2.5-.1.938t-.825.437H13.5v5q0 .425-.288.713T12.5 22h-1ZM12 6q-.825 0-1.413-.588T10 4q0-.825.588-1.413T12 2q.825 0 1.413.588T14 4q0 .825-.588 1.413T12 6Z"/>`),
		g.Group(children),
	)
}