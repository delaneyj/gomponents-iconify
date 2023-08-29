package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LampsOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m8.5 4l.85 3h-4.7l.85-3h3M10 2H4L2 9h10l-2-7m-4 8h2v10h3v2H3v-2h3V10m12.5 0l.85 3h-4.7l.85-3h3M20 8h-6l-2 7h10l-2-7m-4 8h2v4h3v2h-8v-2h3v-4Z"/>`),
		g.Group(children),
	)
}