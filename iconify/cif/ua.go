package cif

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ua(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 301 201"),
		g.Raw(`<g fill="none"><path fill="#005BBB" d="M.5.5h300v200H.5z"/><path fill="#FFD500" d="M.5 100.5h300v100H.5z"/></g>`),
		g.Group(children),
	)
}