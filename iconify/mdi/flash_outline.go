package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlashOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 2h10l-3.5 7H17l-7 13v-8H7V2m2 2v8h3v2.66L14 11h-3.76l3.52-7H9Z"/>`),
		g.Group(children),
	)
}