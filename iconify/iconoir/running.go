package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Running(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M15 7a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm-2.387 1.267l-3.308 4.135l4.135 4.135l-2.067 4.55"/><path d="m6.41 9.508l3.387-3.309l2.816 2.068l2.895 3.308h3.722M8.892 15.71l-1.241.827H4.343"/></g>`),
		g.Group(children),
	)
}