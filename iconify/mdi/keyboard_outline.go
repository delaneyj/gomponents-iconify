package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyboardOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 5a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V7a2 2 0 0 0-2-2H4m0 2h16v10H4V7m1 1v2h2V8H5m3 0v2h2V8H8m3 0v2h2V8h-2m3 0v2h2V8h-2m3 0v2h2V8h-2M5 11v2h2v-2H5m3 0v2h2v-2H8m3 0v2h2v-2h-2m3 0v2h2v-2h-2m3 0v2h2v-2h-2m-9 3v2h8v-2H8Z"/>`),
		g.Group(children),
	)
}