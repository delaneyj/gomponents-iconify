package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cevo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M4.667 8h16l8 13.734L32 16L24 2.135H8zm-1.198 2L0 16l8 13.865h16l3.469-6l-3.333-5.729l-3.469 6h-9.333l-4.667-8.135l3.469-6z"/>`),
		g.Group(children),
	)
}