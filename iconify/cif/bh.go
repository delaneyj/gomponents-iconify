package cif

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bh(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 301 201"),
		g.Raw(`<g fill="none"><path fill="#FFF" d="M.5.5h300v180H.5"/><path fill="#CE1126" d="M300.5.5h-226l46 18l-46 18l46 18l-46 18l46 18l-46 18l46 18l-46 18l46 18l-46 18h226"/></g>`),
		g.Group(children),
	)
}