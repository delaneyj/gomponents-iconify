package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path fill-opacity=".5" d="m16 13.004l-6-.013l-.01 4l6 .013l.01-4Z"/><path d="m19.978 18.002l.026-12l-2-.004l-.026 12l2 .004ZM3.996 10.985l12 .026l.009-4l-12-.026l-.009 4Z"/></g>`),
		g.Group(children),
	)
}