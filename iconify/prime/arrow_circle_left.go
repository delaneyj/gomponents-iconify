package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowCircleLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 21a9 9 0 1 1 9-9a9 9 0 0 1-9 9Zm0-16.5a7.5 7.5 0 1 0 7.5 7.5A7.5 7.5 0 0 0 12 4.5Z"/><path fill="currentColor" d="M12 16.75a.74.74 0 0 1-.53-.22l-4-4a.75.75 0 0 1 0-1.06l4-4a.75.75 0 0 1 1.06 1.06L9.06 12l3.47 3.47a.75.75 0 0 1 0 1.06a.74.74 0 0 1-.53.22Z"/><path fill="currentColor" d="M16 12.75H8a.75.75 0 0 1 0-1.5h8a.75.75 0 0 1 0 1.5Z"/>`),
		g.Group(children),
	)
}