package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarEndOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 9h14v3h2V5c0-1.1-.9-2-2-2h-1V1h-2v2H8V1H6v2H5c-1.1 0-2 .9-2 2v14c0 1.1.9 2 2 2h4v-2H5V9m14-4v2H5V5h14m-3 12h-5v2h5v3l4-4l-4-4v3m4-3v8h2v-8h-2Z"/>`),
		g.Group(children),
	)
}