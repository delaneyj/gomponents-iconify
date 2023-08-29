package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeleteSweepOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 16h4v2h-4v-2m0-8h7v2h-7V8m0 4h6v2h-6v-2m-4-2v8H5v-8h6m2-2H3v10a2 2 0 0 0 2 2h6a2 2 0 0 0 2-2V8m1-3h-3l-1-1H6L5 5H2v2h12V5Z"/>`),
		g.Group(children),
	)
}