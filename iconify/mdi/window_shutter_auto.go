package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WindowShutterAuto(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 14h-2l-3.2 9h1.9l.7-2h3.2l.7 2h1.9L20 14m-2.2 5.7L19 16l1.2 3.7h-2.4M7 9h8v2H7V9m0 3h8v2H7v-2m0 3h8v1.5l-.2.5H7v-2m6.7 5H7v-2h7.5l-.8 2M16 8H6v12H4V8H2V4h18v4h-2v4h-1.4l-.5 1.3l-.1.4V8Z"/>`),
		g.Group(children),
	)
}