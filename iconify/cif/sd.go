package cif

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 301 201"),
		g.Raw(`<g fill="none"><path fill="#D21034" d="M.5.5h300v50H.5z"/><path fill="#FFF" d="M.5 50.5h300v50H.5z"/><path fill="#000" d="M.5 100.5h300v50H.5z"/><path fill="#007229" d="m.5.5l100 75l-100 75z"/></g>`),
		g.Group(children),
	)
}