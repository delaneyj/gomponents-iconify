package cif

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ro(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 301 201"),
		g.Raw(`<g fill="none"><path fill="#002B7F" d="M.5.5h100v200H.5z"/><path fill="#FCD116" d="M100.5.5h100v200h-100z"/><path fill="#CE1126" d="M200.5.5h100v200h-100z"/></g>`),
		g.Group(children),
	)
}