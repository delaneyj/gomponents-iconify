package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RecordStopFortyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M17 15a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V17a2 2 0 0 0-2-2H17Zm27 9c0 11.046-8.954 20-20 20S4 35.046 4 24S12.954 4 24 4s20 8.954 20 20Zm-2.5 0c0-9.665-7.835-17.5-17.5-17.5S6.5 14.335 6.5 24S14.335 41.5 24 41.5S41.5 33.665 41.5 24Z"/>`),
		g.Group(children),
	)
}