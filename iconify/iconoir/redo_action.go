package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RedoAction(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M19 5v6m-3.5-3H9s0 0 0 0s-5 0-5 4.706C4 18 9 18 9 18h8.714"/><path d="M12.5 11.5L16 8l-3.5-3.5"/></g>`),
		g.Group(children),
	)
}