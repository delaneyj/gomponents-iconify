package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextFlowLtr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M1 3h18v2H1zm0 4h14v2H1zm0 4h10v2H1zm0 4h18v2H1z"/>`),
		g.Group(children),
	)
}