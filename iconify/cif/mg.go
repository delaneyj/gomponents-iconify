package cif

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mg(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 301 201"),
		g.Raw(`<g fill="none"><path fill="#FFF" d="M.5.5h300v200H.5z"/><path fill="#FC3D32" d="M100.5.5h200v100h-200z"/><path fill="#007E3A" d="M100.5 100.5h200v100h-200z"/></g>`),
		g.Group(children),
	)
}