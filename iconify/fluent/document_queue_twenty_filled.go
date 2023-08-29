package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentQueueTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 2v4.5A1.5 1.5 0 0 0 11.5 8H16v6.5a1.5 1.5 0 0 1-1.5 1.5h-9A1.5 1.5 0 0 1 4 14.5v-11A1.5 1.5 0 0 1 5.5 2H10ZM6.5 5a.5.5 0 0 0 0 1h2a.5.5 0 0 0 0-1h-2ZM6 7.5a.5.5 0 0 0 .5.5h2a.5.5 0 0 0 0-1h-2a.5.5 0 0 0-.5.5ZM6.5 9a.5.5 0 0 0 0 1h2a.5.5 0 0 0 0-1h-2ZM11 2.25V6.5a.5.5 0 0 0 .5.5h4.25L11 2.25ZM2.5 13a.5.5 0 0 1 .5.5v1A2.5 2.5 0 0 0 5.5 17h9a2.5 2.5 0 0 0 2.5-2.5v-1a.5.5 0 0 1 1 0v1a3.5 3.5 0 0 1-3.5 3.5h-9A3.5 3.5 0 0 1 2 14.5v-1a.5.5 0 0 1 .5-.5Z"/>`),
		g.Group(children),
	)
}