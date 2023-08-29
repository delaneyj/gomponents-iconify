package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Learning(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M6.818 22v-2.857C6.52 16.166 3 14.572 3 10c0-4.57 2.727-8.056 8.182-8c3.927.042 7.636 2.286 7.636 6.858L21 12.286c0 2.286-2.182 2.286-2.182 2.286s.546 5.714-4.364 5.714V22"/><path d="M11 12a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z"/><path stroke-dasharray=".3 2" d="M11 13a3 3 0 1 0 0-6a3 3 0 0 0 0 6Z"/></g>`),
		g.Group(children),
	)
}