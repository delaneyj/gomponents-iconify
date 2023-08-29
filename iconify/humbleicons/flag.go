package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path d="M6 5v15"/><path stroke-linejoin="round" d="M13 5.5V4H7a1 1 0 0 0-1 1v8h7v1.5h7v-9h-7zm0 0v3"/></g>`),
		g.Group(children),
	)
}