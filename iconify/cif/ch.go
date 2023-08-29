package cif

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 301 201"),
		g.Raw(`<g fill="none"><path fill="red" d="M.5.5h300v300H.5z"/><path fill="#FFF" d="M56.75 122.375h187.5v56.25H56.75z"/><path fill="#FFF" d="M122.375 56.75h56.25v187.5h-56.25z"/></g>`),
		g.Group(children),
	)
}