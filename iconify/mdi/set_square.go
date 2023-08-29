package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SetSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m17.7 17.7l-1.1 1.1l-.7-.8l1.1-1l-2-2l-1 1.1l-.7-.7l1.1-1.1l-1.9-1.9l-1.1 1.1l-.7-.7l1.1-1.1l-2-1.9l-1.1 1.1l-.7-.7L9 9L7.1 7.1L6 8.1l-.7-.7l1.1-1.1L4 4v16h16l-2.3-2.3M7 17v-5.8l5.8 5.8H7Z"/>`),
		g.Group(children),
	)
}