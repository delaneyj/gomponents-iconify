package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoveLocation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 22q-4.025-3.425-6.012-6.362T2 10.2q0-3.75 2.413-5.975T10 2q3.175 0 5.588 2.225T18 10.2v.425q0 .225-.05.45q-1.35.275-2.363 1.038T13.95 13.95q-.625 1.075-.787 2.325T13.4 18.8q-.725.775-1.575 1.575T10 22Zm0-7q1.4 0 2.525-.688T14.3 12.5q-.875-.725-1.975-1.113T10 11q-1.225 0-2.325.388T5.7 12.5q.65 1.125 1.775 1.813T10 15Zm0-5q.825 0 1.413-.588T12 8q0-.825-.588-1.413T10 6q-.825 0-1.413.588T8 8q0 .825.588 1.413T10 10Zm9.25 11l-1.4-1.4l1.575-1.6H15.25v-2h4.175l-1.575-1.6l1.4-1.4l4 4l-4 4Z"/>`),
		g.Group(children),
	)
}