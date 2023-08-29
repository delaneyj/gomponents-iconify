package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoodHeart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M21 12a9 9 0 1 0-8.012 8.946M9 10h.01M15 10h.01"/><path d="M9.5 15a3.59 3.59 0 0 0 2.774.99m6.72 5.51l2.518-2.58a1.74 1.74 0 0 0 .004-2.413a1.627 1.627 0 0 0-2.346-.005l-.168.172l-.168-.172a1.627 1.627 0 0 0-2.346-.004a1.74 1.74 0 0 0-.004 2.412l2.51 2.59z"/></g>`),
		g.Group(children),
	)
}