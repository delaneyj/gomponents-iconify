package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GoogleTasks(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.768 5.714a2 2 0 0 1 3.064 2.572L10.833 19.01a2 2 0 1 1-3.064-2.57l8.999-10.726ZM3 12.74a2 2 0 1 1 4 0a2 2 0 0 1-4 0Z"/>`),
		g.Group(children),
	)
}