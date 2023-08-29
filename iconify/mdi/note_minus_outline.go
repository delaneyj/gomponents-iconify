package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoteMinusOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M23 18v2h-8v-2h8m-10 1c0 .7.13 1.37.35 2H5a2 2 0 0 1-2-2V5c0-1.11.89-2 2-2h10l6 6v4.35c-.63-.22-1.3-.35-2-.35v-1h-7V5H5v14h8m1-9h5.5L14 4.5V10Z"/>`),
		g.Group(children),
	)
}