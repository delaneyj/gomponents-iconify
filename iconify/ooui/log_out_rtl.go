package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LogOutRtl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M17 17H9v2h8c1.1 0 2-.9 2-2V3c0-1.1-.9-2-2-2H9v2h8z"/><path fill="currentColor" d="M7 15v-4h8V9H7V5l-6 5z"/>`),
		g.Group(children),
	)
}