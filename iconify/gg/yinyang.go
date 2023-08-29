package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Yinyang(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M14 16a2 2 0 1 1-4 0a2 2 0 0 1 4 0Z"/><path fill-rule="evenodd" d="M22 12c0 5.523-4.477 10-10 10S2 17.523 2 12S6.477 2 12 2s10 4.477 10 10Zm-10 0a4 4 0 0 1 0-8a8 8 0 1 0 0 16a4 4 0 0 0 0-8Zm2-4a2 2 0 1 1-4 0a2 2 0 0 1 4 0Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}