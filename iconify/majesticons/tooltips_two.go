package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TooltipsTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path fill="currentColor" d="M20 10h-8a2 2 0 0 0-2 2v6a2 2 0 0 0 2 2h.93a2 2 0 0 1 1.664.89l.574.862a1 1 0 0 0 1.664 0l.574-.861A2 2 0 0 1 19.07 20H20a2 2 0 0 0 2-2v-6a2 2 0 0 0-2-2z"/><path d="M14 7V4a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v6a2 2 0 0 0 2 2h.93a2 2 0 0 1 1.664.89L7 13.5"/></g>`),
		g.Group(children),
	)
}