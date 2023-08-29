package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Factory(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 18v2h4v-2H4m0-4v2h10v-2H4m6 4v2h4v-2h-4m6-4v2h4v-2h-4m0 4v2h4v-2h-4M2 22V8l5 4V8l5 4V8l5 4l1-10h3l1 10v10H2Z"/>`),
		g.Group(children),
	)
}