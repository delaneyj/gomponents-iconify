package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MultiBubble(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M7.5 22a5.5 5.5 0 1 0-4.764-2.75l-.461 2.475l2.475-.46A5.474 5.474 0 0 0 7.5 22Z"/><path d="M15.282 17.898A7.946 7.946 0 0 0 18 16.93l3.6.67l-.67-3.6A8 8 0 1 0 6.083 8.849"/></g>`),
		g.Group(children),
	)
}