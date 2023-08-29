package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AwardLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="9" r="7"/><path d="M7 14v6.234a1 1 0 0 0 1.514.857l2.972-1.782a1 1 0 0 1 1.028 0l2.972 1.782A1 1 0 0 0 17 20.234V14"/></g>`),
		g.Group(children),
	)
}