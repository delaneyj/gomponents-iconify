package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarSearchOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.3 18.9c.4-.7.7-1.5.7-2.4c0-2.5-2-4.5-4.5-4.5S11 14 11 16.5s2 4.5 4.5 4.5c.9 0 1.7-.2 2.4-.7l3.1 3.1l1.4-1.4l-3.1-3.1m-3.8.1c-1.4 0-2.5-1.1-2.5-2.5s1.1-2.5 2.5-2.5s2.5 1.1 2.5 2.5s-1.1 2.5-2.5 2.5M5 19V9h14v2c.8.5 1.5 1.2 2 2V5c0-1.1-.9-2-2-2h-1V1h-2v2H8V1H6v2H5c-1.1 0-2 .9-2 2v14c0 1.1.9 2 2 2h5.8c-.6-.6-1-1.2-1.3-2H5M19 5v2H5V5h14Z"/>`),
		g.Group(children),
	)
}