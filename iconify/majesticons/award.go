package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Award(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="9" r="7"/><path fill="currentColor" d="M7 20.234V14c.667.667 2.6 2 5 2s4.333-1.333 5-2v6.234a1 1 0 0 1-1.514.857l-2.972-1.782a1 1 0 0 0-1.028 0L8.514 21.09A1 1 0 0 1 7 20.234z"/></g>`),
		g.Group(children),
	)
}