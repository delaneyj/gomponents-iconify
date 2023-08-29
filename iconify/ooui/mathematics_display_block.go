package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MathematicsDisplayBlock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M13 5H5l3 5l-3 5h10v-3h-2v1H9.2l1.8-3l-1.8-3H13v1h2V5zM2 1h16v2H2zm0 16h16v2H2z"/>`),
		g.Group(children),
	)
}