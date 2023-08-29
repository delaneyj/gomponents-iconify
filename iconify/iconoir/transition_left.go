package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TransitionLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M22 18V6a3 3 0 0 0-3-3h-2a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h2a3 3 0 0 0 3-3Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M8 3H6a4 4 0 0 0-4 4v10a4 4 0 0 0 4 4h2m6-9H6m0 0l3-3m-3 3l3 3"/></g>`),
		g.Group(children),
	)
}