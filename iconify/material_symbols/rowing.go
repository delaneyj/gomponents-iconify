package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rowing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m18 24l-3-3v-1.5l-7.1-7.1q-.225.05-.45.075T7 12.5v-2.2q1.25.05 2.55-.537t2.1-1.463l1.4-1.55q.325-.375.763-.562T14.75 6q.95 0 1.6.65t.65 1.6V14q0 .65-.238 1.188t-.662.962l-3.6-3.55v-2.3q-.5.425-1.075.775T10.2 11.7l6.3 6.3H18l3 3l-3 3ZM5.5 20.5L4 19l4.5-4.5L11 17H9l-3.5 3.5ZM15 5q-.825 0-1.413-.588T13 3q0-.825.588-1.413T15 1q.825 0 1.413.588T17 3q0 .825-.588 1.413T15 5Z"/>`),
		g.Group(children),
	)
}