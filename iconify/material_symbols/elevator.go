package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Elevator(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 18h3v-4h1v-2.5q0-.825-.588-1.413T9 9.5H8q-.825 0-1.413.588T6 11.5V14h1v4Zm1.5-9.5q.525 0 .888-.363t.362-.887q0-.525-.363-.888T8.5 6q-.525 0-.888.363t-.362.887q0 .525.363.888T8.5 8.5ZM13 11h5l-2.5-4l-2.5 4Zm2.5 6l2.5-4h-5l2.5 4ZM5 21q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v14q0 .825-.588 1.413T19 21H5Z"/>`),
		g.Group(children),
	)
}