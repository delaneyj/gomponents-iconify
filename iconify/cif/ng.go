package cif

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ng(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 301 201"),
		g.Raw(`<g fill="none"><path fill="#008751" d="M.5.5h300v150H.5z"/><path fill="#FFF" d="M100.5.5h100v150h-100z"/></g>`),
		g.Group(children),
	)
}