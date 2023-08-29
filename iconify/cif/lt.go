package cif

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 301 201"),
		g.Raw(`<g fill="none"><path fill="#FDB913" d="M.5.5h300v180H.5z"/><path fill="#C1272D" d="M.5 90.5h300v90H.5z"/><path fill="#006A44" d="M.5 60.5h300v60H.5z"/></g>`),
		g.Group(children),
	)
}