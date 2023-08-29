package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Html(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16V8l2 5l2-5v8M1 16V8m4 0v8m-4-4h4m2-4h4M9 8v8m11-8v8h3"/>`),
		g.Group(children),
	)
}