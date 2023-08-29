package cif

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 301 201"),
		g.Raw(`<g fill="none"><path fill="#21468B" d="M.5.5h300v200H.5z"/><path fill="#FFF" d="M.5.5h300v133.333H.5z"/><path fill="#AE1C28" d="M.5.5h300v66.667H.5z"/></g>`),
		g.Group(children),
	)
}