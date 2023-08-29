package cif

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 301 201"),
		g.Raw(`<g fill="none"><path fill="#009E49" d="M.5.5h300v180H.5z"/><path fill="#FFF" d="m.5.5l300 90l-300 90z"/><path fill="#FCD116" d="M.5 8.33L274.1 90.5L.5 172.67z"/><path fill="#000" d="m.5.5l150 90l-150 90z"/><path fill="#CE1126" d="m.5 11l132.51 79.5L.5 170z"/></g>`),
		g.Group(children),
	)
}