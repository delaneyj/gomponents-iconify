package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ElevatorOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 14v3q0 .425.288.713T8 18h1q.425 0 .713-.288T10 17v-3q.275-.275.638-.438T11 13v-1.5q0-.825-.588-1.413T9 9.5H8q-.825 0-1.413.588T6 11.5V13q0 .4.363.563T7 14Zm1.5-5.5q.525 0 .888-.363t.362-.887q0-.525-.363-.888T8.5 6q-.525 0-.888.363t-.362.887q0 .525.363.888T8.5 8.5Zm5.4 2.5h3.2q.3 0 .438-.263t-.013-.512l-1.6-2.55q-.15-.25-.425-.25t-.425.25l-1.6 2.55q-.15.25-.012.513T13.9 11Zm2.025 5.325l1.6-2.55q.15-.25.013-.513T17.1 13h-3.2q-.3 0-.438.263t.013.512l1.6 2.55q.15.25.425.25t.425-.25ZM5 21q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v14q0 .825-.588 1.413T19 21H5Zm0-2h14V5H5v14Zm0 0V5v14Z"/>`),
		g.Group(children),
	)
}