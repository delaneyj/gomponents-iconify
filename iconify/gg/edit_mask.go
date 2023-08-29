package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EditMask(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M12 14a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z"/><path fill-rule="evenodd" d="M12 22C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10Zm-2.97-2.57a3 3 0 1 1 5.939 0a8.026 8.026 0 0 0 4.462-4.46a3 3 0 1 1 0-5.938a8.026 8.026 0 0 0-4.462-4.463a3 3 0 1 1-5.939 0a8.026 8.026 0 0 0-4.46 4.462A3.015 3.015 0 0 1 5 9a3 3 0 1 1-.43 5.97a8.026 8.026 0 0 0 4.46 4.46Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}