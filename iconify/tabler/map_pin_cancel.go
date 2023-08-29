package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapPinCancel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M9 11a3 3 0 1 0 6 0a3 3 0 0 0-6 0"/><path d="M12.463 21.431a1.999 1.999 0 0 1-1.876-.531l-4.244-4.243a8 8 0 1 1 13.594-4.655M16 19a3 3 0 1 0 6 0a3 3 0 1 0-6 0m1 2l4-4"/></g>`),
		g.Group(children),
	)
}