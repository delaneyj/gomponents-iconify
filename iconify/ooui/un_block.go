package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UnBlock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M1.22 0L0 1.22l3.06 3.06a9 9 0 0 0 12.66 12.66L18.78 20L20 18.78zM5 11V9h2.78l2 2zm5-10a9 9 0 0 0-4.26 1.08L12.66 9H15v2h-.34l3.26 3.26A9 9 0 0 0 10 1z"/>`),
		g.Group(children),
	)
}