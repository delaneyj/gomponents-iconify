package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SearchLoading(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M8.55 10.55a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm2 1a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm3 0a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/><path fill-rule="evenodd" d="M16.207 4.893a8.001 8.001 0 0 1 .662 10.565c.016.013.03.027.045.042l4.243 4.243a1 1 0 0 1-1.414 1.414L15.5 16.914a.933.933 0 0 1-.042-.045A8.001 8.001 0 0 1 4.893 4.893a8 8 0 0 1 11.314 0Zm-9.9 9.9a6 6 0 1 0 8.486-8.485a6 6 0 0 0-8.485 8.485Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}