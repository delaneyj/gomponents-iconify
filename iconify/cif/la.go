package cif

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func La(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 301 201"),
		g.Raw(`<g fill="none"><path fill="#CE1126" d="M.5.5h300v200H.5z"/><path fill="#002868" d="M.5 50.5h300v100H.5z"/><circle cx="150.5" cy="100.5" r="40" fill="#FFF"/></g>`),
		g.Group(children),
	)
}