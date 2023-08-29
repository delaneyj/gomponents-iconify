package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Candle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M8 5h.926v1.823H8z"/><path d="M10.287 5.973c.11-2.604-.986-3.368-1.61-3.368c-.65 1.177-2.032 1.498-2.329 3.394C5.034 3.567 8.004 3.059 8.153.027c.252-.001 4.403 3.559 2.134 5.946zm1.031 1.436s-.518.644-2.244.644c-1.727 0-3.37-.644-3.37-.644a.677.677 0 0 0-.681.672v7.226c0 .371.305.673.681.673h5.614a.678.678 0 0 0 .682-.673V8.081a.677.677 0 0 0-.682-.672zm-.258 7.607H9.951V8.954h1.109v6.062z"/></g>`),
		g.Group(children),
	)
}