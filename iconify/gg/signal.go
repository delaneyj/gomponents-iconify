package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Signal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 7a1 1 0 1 1 2 0v10a1 1 0 1 1-2 0V7Zm-8 8a1 1 0 1 1 2 0v2a1 1 0 1 1-2 0v-2Zm5-5a1 1 0 0 0-1 1v6a1 1 0 1 0 2 0v-6a1 1 0 0 0-1-1Z"/>`),
		g.Group(children),
	)
}