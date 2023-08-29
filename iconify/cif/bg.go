package cif

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bg(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 301 201"),
		g.Raw(`<g fill="none"><path fill="#FFF" d="M.5.5h300v180H.5z"/><path fill="#00966E" d="M.5 60.5h300v120H.5z"/><path fill="#D62612" d="M.5 120.5h300v60H.5z"/></g>`),
		g.Group(children),
	)
}