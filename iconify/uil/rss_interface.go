package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RssInterface(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 14a1 1 0 0 0 0 2a3 3 0 0 1 3 3a1 1 0 0 0 2 0a5 5 0 0 0-5-5Zm-.71 4.29a1 1 0 1 0 1.42 0a1 1 0 0 0-1.42 0ZM19 4H5a3 3 0 0 0-3 3a1 1 0 0 0 2 0a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1v10a1 1 0 0 1-1 1h-4a1 1 0 0 0 0 2h4a3 3 0 0 0 3-3V7a3 3 0 0 0-3-3ZM3 10a1 1 0 0 0 0 2a7 7 0 0 1 7 7a1 1 0 0 0 2 0a9 9 0 0 0-9-9Z"/>`),
		g.Group(children),
	)
}