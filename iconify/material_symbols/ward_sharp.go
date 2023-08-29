package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WardSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 4H2V2h4v20H4V4Zm3 18V2h13v20H7Zm2-11.475q.45-.275.95-.4T11 10h5q.55 0 1.05.125t.95.4V4H9v6.525ZM13.5 9q-.825 0-1.413-.588T11.5 7q0-.825.588-1.413T13.5 5q.825 0 1.413.588T15.5 7q0 .825-.588 1.413T13.5 9Zm-1 10h2v-2h2v-2h-2v-2h-2v2h-2v2h2v2Z"/>`),
		g.Group(children),
	)
}