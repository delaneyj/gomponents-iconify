package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Facebook(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2.89 2h14.23c.49 0 .88.39.88.88v14.24c0 .48-.39.88-.88.88h-4.08v-6.2h2.08l.31-2.41h-2.39V7.85c0-.7.2-1.18 1.2-1.18h1.28V4.51c-.22-.03-.98-.09-1.86-.09c-1.85 0-3.11 1.12-3.11 3.19v1.78H8.46v2.41h2.09V18H2.89a.89.89 0 0 1-.89-.88V2.88c0-.49.4-.88.89-.88z"/>`),
		g.Group(children),
	)
}