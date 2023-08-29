package cif

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lv(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 301 201"),
		g.Raw(`<g fill="none"><path fill="#9E3039" d="M.5.5h300v150H.5z"/><path fill="#FFF" d="M.5 60.5h300v30H.5z"/></g>`),
		g.Group(children),
	)
}