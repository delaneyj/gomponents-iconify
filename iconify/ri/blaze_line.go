package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlazeLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 9c.667 1.06 1 2.394 1 4c0 3-3.5 4-5 9c-.667-.575-1-1.408-1-2.5c0-3.482 5-5.29 5-10.5Zm-4.5-4a8.309 8.309 0 0 1 1 4c0 5-6 6-4 13C9.833 20.84 9 19.173 9 17c0-3.325 5.5-6 5.5-12ZM10 1c.667 1.333 1 2.833 1 4.5c0 6-9 7.5-3 16.5c-2.5-.5-4.5-3-4.5-6C3.5 9.5 10 8.5 10 1Z"/>`),
		g.Group(children),
	)
}