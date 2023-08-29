package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LiveHelp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 23l-3-3H5q-.825 0-1.413-.588T3 18V4q0-.825.588-1.413T5 2h14q.825 0 1.413.588T21 4v14q0 .825-.588 1.413T19 20h-4l-3 3Zm-.1-6q.525 0 .888-.363t.362-.887q0-.525-.363-.888T11.9 14.5q-.525 0-.888.363t-.362.887q0 .525.363.888T11.9 17Zm-.9-3.85h1.85q0-.425.038-.725t.162-.575q.125-.275.312-.512t.538-.588q.875-.875 1.238-1.463T15.5 7.95q0-1.325-.9-2.138T12.175 5q-1.375 0-2.337.675T8.5 7.55l1.65.65q.175-.675.7-1.087t1.225-.413q.675 0 1.125.363t.45.962q0 .425-.275.9t-.925 1.05q-.425.35-.688.688t-.437.712q-.175.375-.25.788T11 13.15Z"/>`),
		g.Group(children),
	)
}