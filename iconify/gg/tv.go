package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tv(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="m8 6.119l1.413-1.413l2.124 2.124L14.367 4l1.411 1.412l-2.464 2.464H18a2 2 0 0 1 2 2v7a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2v-7a2 2 0 0 1 2-2h3.757L8 6.119Zm10 3.757H6v7h12v-7Z" clip-rule="evenodd"/><path d="M8 19.876h8v1H8v-1Z"/></g>`),
		g.Group(children),
	)
}