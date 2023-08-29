package cif

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cz(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 301 201"),
		g.Raw(`<g fill="none"><path fill="#D7141A" d="M.5.5h300v200H.5z"/><path fill="#FFF" d="M.5.5h300v100H.5z"/><path fill="#11457E" d="M150.5 100.5L.5.5v200z"/></g>`),
		g.Group(children),
	)
}