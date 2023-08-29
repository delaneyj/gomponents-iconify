package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StackedInbox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 17q-.825 0-1.413-.588T5 15V5q0-.825.588-1.413T7 3h14q.825 0 1.413.588T23 5v10q0 .825-.588 1.413T21 17H7Zm7-5q.825 0 1.413-.588T16 10h5V5H7v5h5q0 .825.588 1.413T14 12Zm5 9H3q-.825 0-1.413-.588T1 19V7h2v12h16v2Z"/>`),
		g.Group(children),
	)
}