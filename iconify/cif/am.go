package cif

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Am(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 301 201"),
		g.Raw(`<g fill="none"><path fill="#D90012" d="M.5.5h300v150H.5z"/><path fill="#0033A0" d="M.5 50.5h300v50H.5z"/><path fill="#F2A800" d="M.5 100.5h300v50H.5z"/></g>`),
		g.Group(children),
	)
}