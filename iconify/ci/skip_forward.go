package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SkipForward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 5v14M6 10.571v2.858c0 1.827 0 2.74.384 3.267a2 2 0 0 0 1.413.811c.648.066 1.437-.394 3.016-1.315l2.449-1.428l.008-.005c1.552-.905 2.328-1.358 2.59-1.949a2 2 0 0 0 0-1.62c-.263-.591-1.041-1.045-2.598-1.954l-2.45-1.428c-1.578-.921-2.367-1.381-3.015-1.315a2 2 0 0 0-1.413.812C6 7.83 6 8.745 6 10.57Z"/>`),
		g.Group(children),
	)
}