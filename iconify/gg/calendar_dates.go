package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarDates(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M8 13a1 1 0 1 1 0-2a1 1 0 0 1 0 2Zm0 4a1 1 0 1 1 0-2a1 1 0 0 1 0 2Zm3-1a1 1 0 1 0 2 0a1 1 0 0 0-2 0Zm5 1a1 1 0 1 1 0-2a1 1 0 0 1 0 2Zm-5-5a1 1 0 1 0 2 0a1 1 0 0 0-2 0Zm5 1a1 1 0 1 1 0-2a1 1 0 0 1 0 2ZM8 7a1 1 0 0 0 0 2h8a1 1 0 1 0 0-2H8Z"/><path fill-rule="evenodd" d="M6 3a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3V6a3 3 0 0 0-3-3H6Zm12 2H6a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1V6a1 1 0 0 0-1-1Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}