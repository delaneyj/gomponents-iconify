package akar_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cake(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-linejoin="round" d="M3 13a4 4 0 0 1 4-4h10a4 4 0 0 1 4 4v7a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-7Z"/><path d="m3 13l2.914 2.331c1.187.95 2.9.855 3.975-.22v0a2.985 2.985 0 0 1 4.222 0v0a2.985 2.985 0 0 0 3.975.22L21 13"/><path stroke-linejoin="round" d="M12 6a2 2 0 0 1-2-2c0-.876.677-1.576 1.273-2.217L12 1l.727.783C13.323 2.424 14 3.124 14 4a2 2 0 0 1-2 2Z"/></g>`),
		g.Group(children),
	)
}