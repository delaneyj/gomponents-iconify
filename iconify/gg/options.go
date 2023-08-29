package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Options(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7 3a4.002 4.002 0 0 1 3.874 3H19v2h-8.126A4.002 4.002 0 0 1 3 7a4 4 0 0 1 4-4Zm0 6a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm10 11a4.002 4.002 0 0 1-3.874-3H5v-2h8.126A4.002 4.002 0 0 1 21 16a4 4 0 0 1-4 4Zm0-2a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}