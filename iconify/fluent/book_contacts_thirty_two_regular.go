package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookContactsThirtyTwoRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M21 16.5a1.5 1.5 0 0 0-1.5-1.5h-7a1.5 1.5 0 0 0-1.5 1.5v.5c0 1.971 1.86 4 5 4c3.14 0 5-2.029 5-4v-.5Zm-2.25-5.25a2.75 2.75 0 1 0-5.5 0a2.75 2.75 0 0 0 5.5 0ZM25 26H7.085A1.5 1.5 0 0 0 8.5 27H26a1 1 0 1 1 0 2H8.5A3.5 3.5 0 0 1 5 25.5v-19A3.5 3.5 0 0 1 8.5 3h15A3.5 3.5 0 0 1 27 6.5V24a2 2 0 0 1-2 2ZM7 6.5V24h18V6.5A1.5 1.5 0 0 0 23.5 5h-15A1.5 1.5 0 0 0 7 6.5Z"/>`),
		g.Group(children),
	)
}