package cif

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 301 201"),
		g.Raw(`<g fill="none"><path fill="#377E3F" d="M.5.5h300v200H.5z"/><path fill="#FFF" d="M.5 40.5h300v120H.5z"/><path fill="#B40A2D" d="M.5 60.5h300v80H.5z"/><path fill="#ECC81D" d="m150.5 64.319l23.511 72.361l-61.554-44.721h76.085l-61.554 44.721z"/></g>`),
		g.Group(children),
	)
}