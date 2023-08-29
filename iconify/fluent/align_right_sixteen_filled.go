package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignRightSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M14 1.5a.5.5 0 0 0-1 0v13a.5.5 0 0 0 1 0v-13Zm-3.75.5c.966 0 1.75.784 1.75 1.75v1.5A1.75 1.75 0 0 1 10.25 7h-6.5A1.75 1.75 0 0 1 2 5.25v-1.5C2 2.784 2.784 2 3.75 2h6.5Zm0 7c.966 0 1.75.784 1.75 1.75v1.5A1.75 1.75 0 0 1 10.25 14h-4.5A1.75 1.75 0 0 1 4 12.25v-1.5C4 9.784 4.784 9 5.75 9h4.5Z"/>`),
		g.Group(children),
	)
}