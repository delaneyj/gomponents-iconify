package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewDetailsRtl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 8V6h9v2zm9-3H1V3h11zm1-2h6v6h-6zm-1 13H3v-2h9zm0-3H1v-2h11zm1-2h6v6h-6z"/>`),
		g.Group(children),
	)
}