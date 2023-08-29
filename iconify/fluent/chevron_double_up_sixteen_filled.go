package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronDoubleUpSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4.26 8.3a.75.75 0 1 1-1.02-1.1l4.25-4a.75.75 0 0 1 1.02 0l4.25 4a.75.75 0 1 1-1.02 1.1L8 4.773L4.26 8.3Zm0 4a.75.75 0 0 1-1.02-1.1l4.25-4a.75.75 0 0 1 1.02 0l4.25 4a.75.75 0 1 1-1.02 1.1L8 8.773L4.26 12.3Z"/>`),
		g.Group(children),
	)
}