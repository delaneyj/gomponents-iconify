package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StackedEmail(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21q-.825 0-1.413-.588T1 19V6.5h2V19h16.5v2H3Zm4-4q-.825 0-1.413-.588T5 15V5q0-.825.588-1.413T7 3h14q.825 0 1.413.588T23 5v10q0 .825-.588 1.413T21 17H7Zm7-4.7l7-4.875V5l-7 4.85L7 5v2.425l7 4.875Z"/>`),
		g.Group(children),
	)
}