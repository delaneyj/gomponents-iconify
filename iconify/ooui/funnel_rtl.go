package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FunnelRtl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M1 1h18l-9 12z"/><path fill="currentColor" d="m8 19l4-2V9H8z"/>`),
		g.Group(children),
	)
}