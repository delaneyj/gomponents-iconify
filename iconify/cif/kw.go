package cif

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Kw(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 301 201"),
		g.Raw(`<g fill="none"><path fill="#007A3D" d="M.5.5h300v50H.5z"/><path fill="#FFF" d="M.5 50.5h300v50H.5z"/><path fill="#CE1126" d="M.5 100.5h300v50H.5z"/><path fill="#000" d="m.5.5l75 50v50l-75 50z"/></g>`),
		g.Group(children),
	)
}