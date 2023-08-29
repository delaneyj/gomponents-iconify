package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImageLockLtr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M16 17H2l3.5-4.5l2.5 3l3-3.81A2 2 0 0 1 10 10V6H2a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-6h-5.75z"/><path fill="currentColor" d="M19 4h-.5V2.5a2.5 2.5 0 0 0-5 0V4H13a1 1 0 0 0-1 1v4a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1V5a1 1 0 0 0-1-1zm-3 4a1 1 0 1 1 1-1a1 1 0 0 1-1 1zm1.5-4h-3V2.75C14.5 2 14.5 1 16 1s1.5 1 1.5 1.75z"/>`),
		g.Group(children),
	)
}