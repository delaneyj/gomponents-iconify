package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DockTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10.53 3.22a.75.75 0 0 0-1.06 0l-2.5 2.5a.75.75 0 0 0 1.06 1.06l1.22-1.22v6.69a.75.75 0 0 0 1.5 0V5.56l1.22 1.22a.75.75 0 1 0 1.06-1.06l-2.5-2.5ZM3.5 10.5a1 1 0 0 1 1-1h2.75a.75.75 0 0 0 0-1.5H4.5A2.5 2.5 0 0 0 2 10.5v3A2.5 2.5 0 0 0 4.5 16h11a2.5 2.5 0 0 0 2.5-2.5v-3A2.5 2.5 0 0 0 15.5 8h-2.75a.75.75 0 0 0 0 1.5h2.75a1 1 0 0 1 1 1v3a1 1 0 0 1-1 1h-11a1 1 0 0 1-1-1v-3Z"/>`),
		g.Group(children),
	)
}