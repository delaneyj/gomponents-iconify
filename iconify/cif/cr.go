package cif

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 301 201"),
		g.Raw(`<g fill="none"><path fill="#002B7F" d="M.5.5h300v180H.5z"/><path fill="#FFF" d="M.5 30.5h300v120H.5z"/><path fill="#CE1126" d="M.5 60.5h300v60H.5z"/></g>`),
		g.Group(children),
	)
}