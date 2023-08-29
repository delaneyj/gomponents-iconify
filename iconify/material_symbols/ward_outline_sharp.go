package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WardOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 4H2V2h4v20H4V4Zm3 18V2h13v20H7Zm2-11.475q.45-.275.95-.4T11 10h5q.55 0 1.05.125t.95.4V4H9v6.525ZM13.5 9q-.825 0-1.413-.588T11.5 7q0-.825.588-1.413T13.5 5q.825 0 1.413.588T15.5 7q0 .825-.588 1.413T13.5 9ZM9 20h9v-6q0-.825-.588-1.413T16 12h-5q-.825 0-1.413.588T9 14v6Zm3.5-1h2v-2h2v-2h-2v-2h-2v2h-2v2h2v2ZM9 20h9h-9Z"/>`),
		g.Group(children),
	)
}