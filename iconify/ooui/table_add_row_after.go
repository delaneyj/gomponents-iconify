package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableAddRowAfter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 0v8H2v12h16V8h-1V0zm8 10v3h3v2h-3v3H9v-3H6v-2h3v-3zm4-4v2H5V6zm0-4v2H5V2z"/>`),
		g.Group(children),
	)
}