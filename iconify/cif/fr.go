package cif

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 301 201"),
		g.Raw(`<g fill="none"><path fill="#ED2939" d="M.5.5h300v200H.5z"/><path fill="#FFF" d="M.5.5h200v200H.5z"/><path fill="#002395" d="M.5.5h100v200H.5z"/></g>`),
		g.Group(children),
	)
}