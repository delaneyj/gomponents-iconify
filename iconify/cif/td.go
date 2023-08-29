package cif

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Td(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 301 201"),
		g.Raw(`<g fill="none"><path fill="#002664" d="M.5.5h300v200H.5z"/><path fill="#FECB00" d="M100.5.5h200v200h-200z"/><path fill="#C60C30" d="M200.5.5h100v200h-100z"/></g>`),
		g.Group(children),
	)
}