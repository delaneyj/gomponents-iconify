package cif

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ga(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 301 201"),
		g.Raw(`<g fill="none"><path fill="#3A75C4" d="M.5 0h300v225H.5z"/><path fill="#FCD116" d="M.5 0h300v150H.5z"/><path fill="#009E60" d="M.5 0h300v75H.5z"/></g>`),
		g.Group(children),
	)
}