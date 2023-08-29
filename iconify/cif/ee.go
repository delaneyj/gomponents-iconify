package cif

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ee(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 301 201"),
		g.Raw(`<g fill="none"><path fill="#4891D9" d="M.5.5h300v190H.5z"/><path fill="#000" d="M.5 63.833h300V190.5H.5z"/><path fill="#FFF" d="M.5 127.166h300v63.333H.5z"/></g>`),
		g.Group(children),
	)
}