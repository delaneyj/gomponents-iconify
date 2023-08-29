package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CopyRtl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M17 3H9v2H7V3c0-1.1.895-2 2-2h8c1.1 0 2 .895 2 2v8c0 1.1-.895 2-2 2h-2v-2h2z"/><path fill="currentColor" d="M3 9v8h8V9zm8-2c1.1 0 2 .895 2 2v8c0 1.1-.895 2-2 2H3c-1.1 0-2-.895-2-2V9c0-1.1.895-2 2-2z"/>`),
		g.Group(children),
	)
}