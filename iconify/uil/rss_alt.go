package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RssAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2.88 16.88a3 3 0 1 0 4.24 4.24a3 3 0 0 0 0-4.24a3.08 3.08 0 0 0-4.24 0Zm2.83 2.83a1 1 0 0 1-1.42 0a1 1 0 0 1 0-1.42a1 1 0 0 1 1.42 0a1 1 0 0 1 0 1.42ZM5 12a1 1 0 0 0 0 2a5 5 0 0 1 5 5a1 1 0 0 0 2 0a7 7 0 0 0-7-7Zm0-4a1 1 0 0 0 0 2a9 9 0 0 1 9 9a1 1 0 0 0 2 0A11 11 0 0 0 5 8Z"/>`),
		g.Group(children),
	)
}