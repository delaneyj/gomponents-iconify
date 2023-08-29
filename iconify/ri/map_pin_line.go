package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapPinLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 20.9l4.95-4.95a7 7 0 1 0-9.9 0L12 20.9Zm0 2.828l-6.364-6.364a9 9 0 1 1 12.728 0L12 23.728ZM12 13a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm0 2a4 4 0 1 1 0-8a4 4 0 0 1 0 8Z"/>`),
		g.Group(children),
	)
}