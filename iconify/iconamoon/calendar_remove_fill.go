package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarRemoveFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M17 3a1 1 0 1 0-2 0H9a1 1 0 0 0-2 0H4a1 1 0 0 0-1 1v14a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3V4a1 1 0 0 0-1-1h-3ZM8 6a1 1 0 0 1-1-1H5v2h14V5h-2a1 1 0 1 1-2 0H9a1 1 0 0 1-1 1Zm2 7a1 1 0 1 0 0 2h4a1 1 0 1 0 0-2h-4Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}