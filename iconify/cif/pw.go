package cif

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pw(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 301 201"),
		g.Raw(`<g fill="none"><path fill="#4AADD6" d="M.5.75h300v187.5H.5z"/><circle cx="131.75" cy="94.5" r="56.25" fill="#FFDE00"/></g>`),
		g.Group(children),
	)
}