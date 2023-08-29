package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartBar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M2 2h2v16h16v2H2V2m4 14V8h4v8H6m5 0V4h4v12h-4m5 0v-5h4v5h-4Z"/>`),
		g.Group(children),
	)
}