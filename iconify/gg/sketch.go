package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sketch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M5.209 3h13.694l1.209 7.253l-8.056 10.933L4 10.253L5.209 3Zm1.694 2l-.791 4.747l5.944 8.067L18 9.747L17.209 5H6.903Z" clip-rule="evenodd"/><path d="M8.056 8h8v2h-8V8Z"/></g>`),
		g.Group(children),
	)
}