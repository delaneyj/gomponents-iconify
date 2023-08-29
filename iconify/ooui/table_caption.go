package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableCaption(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 8a2 2 0 0 0-2 2v6a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-6a2 2 0 0 0-2-2zm0 2h7v2H2zm0 6v-2h7v2zm16 0h-7v-2h7zm0-4h-7v-2h7zM2 2h16v4H2z"/>`),
		g.Group(children),
	)
}