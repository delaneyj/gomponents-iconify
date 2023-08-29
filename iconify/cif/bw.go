package cif

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bw(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 301 201"),
		g.Raw(`<g fill="none"><path fill="#75AADB" d="M.5.5h300v200H.5z"/><path fill="#FFF" d="M.5 75.5h300v50H.5z"/><path fill="#000" d="M.5 83.833h300v33.333H.5z"/></g>`),
		g.Group(children),
	)
}