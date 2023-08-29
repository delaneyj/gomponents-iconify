package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableAddRowBefore(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M17 20v-8h1V0H2v12h1v8zM9 10V7H6V5h3V2h2v3h3v2h-3v3zm-4 4v-2h10v2zm0 4v-2h10v2z"/>`),
		g.Group(children),
	)
}