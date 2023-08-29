package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ToggleOn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M7 15a3 3 0 1 1 0-6a3 3 0 0 1 0 6Z"/><path fill-rule="evenodd" d="M24 12a7 7 0 0 0-7-7H7a7 7 0 0 0 0 14h10a7 7 0 0 0 7-7Zm-7-5H7a5 5 0 0 0 0 10h10a5 5 0 0 0 0-10Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}