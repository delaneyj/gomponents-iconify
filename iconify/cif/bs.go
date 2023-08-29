package cif

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bs(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 301 201"),
		g.Raw(`<g fill="none"><path fill="#00ABC9" d="M.5.5h300v150H.5z"/><path fill="#FAE042" d="M.5 50.5h300v50H.5z"/><path fill="#000" d="M.5.5v150l125-75z"/></g>`),
		g.Group(children),
	)
}