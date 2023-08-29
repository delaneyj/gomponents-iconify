package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImageAddLtr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M16 17H2l3.5-4.5l2.5 3l3.5-4.5l.5.67V8H8V6H2a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-6h-5.75z"/><path fill="currentColor" d="M16 4V0h-2v4h-4v2h4v4h2V6h4V4z"/>`),
		g.Group(children),
	)
}