package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 17.75a.74.74 0 0 1-.53-.22l-5-5a.75.75 0 0 1 0-1.06l5-5a.75.75 0 0 1 1.06 1.06L10.06 12l4.47 4.47a.75.75 0 0 1 0 1.06a.74.74 0 0 1-.53.22Z"/>`),
		g.Group(children),
	)
}