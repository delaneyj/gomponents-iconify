package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dribbble(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M2 12c0 5.523 4.477 10 10 10s10-4.477 10-10S17.523 2 12 2S2 6.477 2 12Z"/><path d="M16.673 20.843C15.5 14 12.5 8 8.5 2.63"/><path d="M2.067 10.84C6 11 15.283 10.5 19.142 5m2.826 7.81C15.344 10.84 7.5 14 5.23 19.361"/></g>`),
		g.Group(children),
	)
}