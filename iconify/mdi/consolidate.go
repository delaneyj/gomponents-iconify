package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Consolidate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 9h2V4h2v5a2 2 0 0 1-2 2h-2v2l-3-3l3-3m-4 3a2 2 0 1 0-2 2a2 2 0 0 0 2-2M2 11v5h2v-5h2v2l3-3l-3-3v2H4a2 2 0 0 0-2 2m13 5l-3-3l-3 3h2v2a2 2 0 0 0 2 2h5v-2h-5v-2"/>`),
		g.Group(children),
	)
}