package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NotificationOffFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M10 3a2 2 0 1 1 3.98.284A7.003 7.003 0 0 1 19 10v2.343a1 1 0 0 1-2 0V10a5 5 0 0 0-6.668-4.715a1 1 0 0 1-.667-1.886a7 7 0 0 1 .355-.115A2.018 2.018 0 0 1 10 3ZM7.604 6.19l11.09 11.09l.026.026l.981.98a.679.679 0 0 1 .012.013l.994.994a1 1 0 0 1-1.414 1.414L18.586 20H4a1 1 0 1 1 0-2h1v-8a6.97 6.97 0 0 1 .646-2.94L3.293 4.707a1 1 0 0 1 1.414-1.414L7.564 6.15a.982.982 0 0 1 .04.04ZM13 23a1 1 0 1 0 0-2h-2a1 1 0 1 0 0 2h2Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}