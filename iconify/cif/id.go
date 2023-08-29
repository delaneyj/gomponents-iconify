package cif

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Id(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 301 201"),
		g.Raw(`<g fill="none"><path fill="#CE1126" d="M.5.5h300v200H.5z"/><path fill="#FFF" d="M.5 100.5h300v100H.5z"/></g>`),
		g.Group(children),
	)
}