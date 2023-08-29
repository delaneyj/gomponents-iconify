package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarOneFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M16 2a1 1 0 0 1 1 1h3a1 1 0 0 1 1 1v14a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3V4a1 1 0 0 1 1-1h3a1 1 0 0 1 2 0h6a1 1 0 0 1 1-1ZM7 5a1 1 0 0 0 2 0h6a1 1 0 1 0 2 0h2v2H5V5h2Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}