package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Trolley(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 17V5H2V3h4v12h15v2H4Zm2 5q-.825 0-1.413-.588T4 20q0-.825.588-1.413T6 18q.825 0 1.413.588T8 20q0 .825-.588 1.413T6 22Zm1-8V8h6v6H7Zm7 0V8h6v6h-6Zm5 8q-.825 0-1.413-.588T17 20q0-.825.588-1.413T19 18q.825 0 1.413.588T21 20q0 .825-.588 1.413T19 22Z"/>`),
		g.Group(children),
	)
}