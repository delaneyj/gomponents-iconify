package cif

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mv(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 301 201"),
		g.Raw(`<g fill="none"><path fill="#D21034" d="M.5.5h300v200H.5z"/><path fill="#007E3A" d="M50.5 50.5h200v100h-200z"/><circle cx="163" cy="100.5" r="33.333" fill="#FFF"/><circle cx="175.5" cy="100.5" r="33.333" fill="#007E3A"/></g>`),
		g.Group(children),
	)
}