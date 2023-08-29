package cif

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lu(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 301 201"),
		g.Raw(`<g fill="none"><path fill="#00A1DE" d="M.5 90.5h300v90H.5z"/><path fill="#ED2939" d="M.5.5h300v90H.5z"/><path fill="#FFF" d="M.5 60.5h300v60H.5z"/></g>`),
		g.Group(children),
	)
}