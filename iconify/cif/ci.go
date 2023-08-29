package cif

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ci(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 301 201"),
		g.Raw(`<g fill="none"><path fill="#FFF" d="M.5.5h300v200H.5z"/><path fill="#F77F00" d="M.5.5h100v200H.5z"/><path fill="#009E60" d="M200.5.5h100v200h-100z"/></g>`),
		g.Group(children),
	)
}