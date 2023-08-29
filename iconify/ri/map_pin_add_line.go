package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapPinAddLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 20.9l4.95-4.95a7 7 0 1 0-9.9 0L12 20.9Zm0 2.828l-6.364-6.364a9 9 0 1 1 12.728 0L12 23.728ZM11 10V7h2v3h3v2h-3v3h-2v-3H8v-2h3Z"/>`),
		g.Group(children),
	)
}