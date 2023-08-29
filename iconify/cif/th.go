package cif

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Th(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 301 201"),
		g.Raw(`<g fill="none"><path fill="#ED1C24" d="M.5.5h300v200H.5z"/><path fill="#FFF" d="M.5 33.833h300v133.333H.5z"/><path fill="#241D4F" d="M.5 67.166h300v66.667H.5z"/></g>`),
		g.Group(children),
	)
}