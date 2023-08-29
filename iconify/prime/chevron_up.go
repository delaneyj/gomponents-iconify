package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 15.25a.74.74 0 0 1-.53-.22L12 10.56L7.53 15a.75.75 0 0 1-1.06-1l5-5a.75.75 0 0 1 1.06 0l5 5a.75.75 0 0 1 0 1.06a.74.74 0 0 1-.53.19Z"/>`),
		g.Group(children),
	)
}