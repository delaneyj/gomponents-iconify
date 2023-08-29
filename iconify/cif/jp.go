package cif

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Jp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 301 201"),
		g.Raw(`<g fill="none"><path fill="#FFF" d="M.5.5h300v200H.5z"/><circle cx="150.5" cy="100.5" r="60" fill="#BC002D"/></g>`),
		g.Group(children),
	)
}