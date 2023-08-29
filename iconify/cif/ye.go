package cif

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ye(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 301 201"),
		g.Raw(`<g fill="none"><path fill="#FFF" d="M0 0h300v200H0z"/><path fill="#CE1126" d="M0 0h300v66.667H0z"/><path fill="#000" d="M0 133.333h300V200H0z"/></g>`),
		g.Group(children),
	)
}