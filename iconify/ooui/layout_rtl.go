package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LayoutRtl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M1 12v7h18V1h-7v11z"/><path fill="currentColor" d="M1 1v8h8V1zm2 2h4v4H3z"/>`),
		g.Group(children),
	)
}