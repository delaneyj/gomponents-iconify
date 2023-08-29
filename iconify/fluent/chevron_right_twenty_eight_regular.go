package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronRightTwentyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M9.97 4.22a.75.75 0 0 1 1.06 0l8.75 8.75a.75.75 0 0 1 0 1.06l-8.75 8.75a.75.75 0 1 1-1.06-1.06l8.22-8.22l-8.22-8.22a.75.75 0 0 1 0-1.06Z"/>`),
		g.Group(children),
	)
}