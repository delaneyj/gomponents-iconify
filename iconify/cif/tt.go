package cif

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 301 201"),
		g.Raw(`<g fill="none"><path fill="#CE1126" d="M.5.5h300v180H.5z"/><path fill="#FFF" d="m.5.5l208.248 180H300.5L92.253.5z"/><path fill="#000" d="m15.792.5l208.247 180h61.169L76.961.5z"/></g>`),
		g.Group(children),
	)
}