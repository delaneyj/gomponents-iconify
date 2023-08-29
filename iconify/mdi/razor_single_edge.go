package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RazorSingleEdge(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 5v3H2V5h20m0 8v5H2v-5h1.5c.83 0 1.5-.67 1.5-1.5S4.33 10 3.5 10H2V9h20v1h-1.5c-.83 0-1.5.67-1.5 1.5s.67 1.5 1.5 1.5H22m-9-2c0-.55-.45-1-1-1s-1 .45-1 1v2c0 .55.45 1 1 1s1-.45 1-1v-2Z"/>`),
		g.Group(children),
	)
}