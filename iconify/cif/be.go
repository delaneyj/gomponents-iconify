package cif

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Be(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 301 201"),
		g.Raw(`<g fill="none"><path fill="#000" d="M.5.5h300v260H.5z"/><path fill="#FAE042" d="M100.5.5h100v260h-100z"/><path fill="#ED2939" d="M200.5.5h100v260h-100z"/></g>`),
		g.Group(children),
	)
}