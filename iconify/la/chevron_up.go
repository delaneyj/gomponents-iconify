package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 5.688L1.594 20.093l.687.718l3.906 3.907l.72.687L16 16.313l9.094 9.093l.718-.687l3.907-3.907l.687-.718zm0 2.843l11.563 11.594l-2.438 2.438l-8.406-8.375L16 13.5l-.719.688l-8.406 8.374l-2.438-2.437z"/>`),
		g.Group(children),
	)
}