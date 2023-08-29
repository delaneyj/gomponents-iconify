package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QuestionMark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.6 16q0-2.025.363-2.913T12.5 11.15q1.025-.9 1.563-1.563t.537-1.512q0-1.025-.687-1.7T12 5.7q-1.275 0-1.938.775T9.126 8.05L6.55 6.95q.525-1.6 1.925-2.775T12 3q2.625 0 4.038 1.463t1.412 3.512q0 1.25-.537 2.138t-1.688 2.012Q14 13.3 13.738 13.913T13.475 16H10.6Zm1.4 6q-.825 0-1.412-.588T10 20q0-.825.588-1.413T12 18q.825 0 1.413.588T14 20q0 .825-.588 1.413T12 22Z"/>`),
		g.Group(children),
	)
}