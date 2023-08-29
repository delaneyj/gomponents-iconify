package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HelpCenterRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 18q.525 0 .888-.363t.362-.887q0-.525-.363-.888T12 15.5q-.525 0-.888.363t-.362.887q0 .525.363.888T12 18Zm3.6-9.05q0-1.35-.913-2.15T12.276 6q-1.125 0-1.988.463T8.925 7.8q-.175.275-.013.6t.513.475q.3.125.625.013t.55-.413q.275-.375.688-.575t.887-.2q.7 0 1.137.375t.438.95q0 .45-.3.95t-.9 1q-.65.575-.975 1.075t-.425 1.175q-.05.35.212.637t.663.288q.35 0 .638-.238t.362-.637q.075-.4.275-.713t.7-.812q.95-.95 1.275-1.525T15.6 8.95ZM5 21q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v14q0 .825-.588 1.413T19 21H5Z"/>`),
		g.Group(children),
	)
}