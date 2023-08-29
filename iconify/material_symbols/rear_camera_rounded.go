package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RearCameraRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.5 8q.425 0 .713-.288T17.5 7q0-.425-.288-.713T16.5 6q-.425 0-.713.288T15.5 7q0 .425.288.713T16.5 8ZM13 19h7V5h-7v14Zm-9 2q-.825 0-1.413-.588T2 19v-6h4.2l-.9.9q-.275.275-.275.7t.275.7q.275.275.7.275t.7-.275l2.6-2.6q.3-.3.3-.7t-.3-.7L6.7 8.7q-.275-.275-.7-.275t-.7.275q-.275.275-.275.7t.275.7l.9.9H2V5q0-.825.588-1.413T4 3h16q.825 0 1.413.588T22 5v14q0 .825-.588 1.413T20 21H4Z"/>`),
		g.Group(children),
	)
}