package cif

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 301 201"),
		g.Raw(`<g fill="none"><path fill="#FFF" d="M.5.75h300v187.5H.5z"/><path fill="#DC143C" d="M.5 94.5h300v93.75H.5z"/></g>`),
		g.Group(children),
	)
}