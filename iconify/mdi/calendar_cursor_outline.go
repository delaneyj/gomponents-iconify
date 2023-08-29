package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarCursorOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 9h14v2.1l2 1.7V5c0-1.1-.9-2-2-2h-1V1h-2v2H8V1H6v2H5c-1.1 0-2 .9-2 2v14c0 1.1.9 2 2 2h8v-2H5V9m14-4v2H5V5h14M7 16v-5h5v5H7m15.9 1.7c-.1.1-.2.2-.3.2l-1.9.4l1.3 2.8c.2.3 0 .6-.3.8l-2.1 1c-.2.1-.2.1-.3.1c-.2 0-.4-.1-.5-.3l-1.3-2.8L16 21c-.1.1-.2.1-.4.1c-.3 0-.6-.3-.6-.6v-9c0-.3.3-.6.6-.6c.1 0 .3.1.4.1l6.8 5.8c.2.3.3.7.1.9Z"/>`),
		g.Group(children),
	)
}