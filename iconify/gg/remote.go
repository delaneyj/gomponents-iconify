package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Remote(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m17.051 4.322l1.415 1.414l-4.243 4.243l4.243 4.242l-1.415 1.415l-5.656-5.657l5.656-5.657ZM6.949 19.678l-1.415-1.414l4.243-4.242l-4.243-4.243L6.95 8.365l5.656 5.657l-5.656 5.656Z"/>`),
		g.Group(children),
	)
}