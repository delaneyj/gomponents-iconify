package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignLeftSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 1.5a.5.5 0 0 1 1 0v13a.5.5 0 0 1-1 0v-13Zm3.75.5A1.75 1.75 0 0 0 4 3.75v1.5C4 6.216 4.784 7 5.75 7h6.5A1.75 1.75 0 0 0 14 5.25v-1.5A1.75 1.75 0 0 0 12.25 2h-6.5Zm0 7A1.75 1.75 0 0 0 4 10.75v1.5c0 .966.784 1.75 1.75 1.75h4.5A1.75 1.75 0 0 0 12 12.25v-1.5A1.75 1.75 0 0 0 10.25 9h-4.5Z"/>`),
		g.Group(children),
	)
}