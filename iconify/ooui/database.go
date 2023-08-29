package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Database(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4 1h12l2 2v2H2V3zM2 8h16v4H2zm16 9v-2H2v2l2 2h12z"/>`),
		g.Group(children),
	)
}