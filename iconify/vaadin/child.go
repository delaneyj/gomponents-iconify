package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Child(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M10 5a2 2 0 1 1-3.999.001A2 2 0 0 1 10 5z"/><path fill="currentColor" d="m12.79 10.32l-2.6-2.63A2.311 2.311 0 0 0 8.54 7H7.469c-.648 0-1.235.264-1.659.69l-2.6 2.63a.73.73 0 1 0 .998 1.003L6 9.53V16h1.5v-4h1v4H10V9.53l1.75 1.8a.73.73 0 1 0 1.041-1.009z"/>`),
		g.Group(children),
	)
}