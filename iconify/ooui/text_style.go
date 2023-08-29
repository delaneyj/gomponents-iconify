package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextStyle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 17h16v2H2zm9.34-15h3.31l2 14h-3.19l-.29-2.88H8L6.43 16H3.37zm-2 8.71h3.55l-.61-5.51z"/>`),
		g.Group(children),
	)
}