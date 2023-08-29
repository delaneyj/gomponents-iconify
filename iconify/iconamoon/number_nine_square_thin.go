package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberNineSquareThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2v-14Z"/><path fill="currentColor" d="M10.582 16.726a.5.5 0 1 0 .836.548l-.836-.548Zm4.361-4.831a.5.5 0 1 0-.836-.548l.836.548ZM9.5 10A2.5 2.5 0 0 1 12 7.5v-1A3.5 3.5 0 0 0 8.5 10h1ZM12 7.5a2.5 2.5 0 0 1 2.5 2.5h1A3.5 3.5 0 0 0 12 6.5v1Zm2.5 2.5a2.5 2.5 0 0 1-2.5 2.5v1a3.5 3.5 0 0 0 3.5-3.5h-1ZM12 12.5A2.5 2.5 0 0 1 9.5 10h-1a3.5 3.5 0 0 0 3.5 3.5v-1Zm-.582 4.774l3.525-5.38l-.836-.547l-3.525 5.379l.836.548Z"/></g>`),
		g.Group(children),
	)
}