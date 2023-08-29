package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MenuBurgerVerticalFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M19 3.5a1 1 0 1 0-2 0v17a1 1 0 1 0 2 0v-17Zm-7-1a1 1 0 0 1 1 1v17a1 1 0 1 1-2 0v-17a1 1 0 0 1 1-1Zm-6 0a1 1 0 0 1 1 1v17a1 1 0 1 1-2 0v-17a1 1 0 0 1 1-1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}