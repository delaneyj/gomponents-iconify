package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Images(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 2H8c-1.103 0-2 .897-2 2v12c0 1.103.897 2 2 2h12c1.103 0 2-.897 2-2V4c0-1.103-.897-2-2-2zM8 16V4h12l.002 12H8z"/><path fill="currentColor" d="M4 8H2v12c0 1.103.897 2 2 2h12v-2H4V8z"/><path fill="currentColor" d="m12 12l-1-1l-2 3h10l-4-6z"/>`),
		g.Group(children),
	)
}