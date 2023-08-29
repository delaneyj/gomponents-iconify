package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmartHomeHeat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M17 11H7a1 1 0 1 0 0 2h3v2H7a3 3 0 1 1 0-6h10a3 3 0 1 1 0 6h-3v-2h3a1 1 0 1 0 0-2Z"/><path fill-rule="evenodd" d="M0 12a7 7 0 0 1 7-7h10a7 7 0 1 1 0 14H7a7 7 0 0 1-7-7Zm7-5h10a5 5 0 0 1 0 10H7A5 5 0 0 1 7 7Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}