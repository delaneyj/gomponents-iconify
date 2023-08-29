package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudComputing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 20h-3a1 1 0 0 1-1-1v-3a5 5 0 0 0 1.42-9.78A7 7 0 0 0 5.06 8.11A4 4 0 0 0 6 16h1v3a1 1 0 0 1-1 1H3a1 1 0 0 0 0 2h3a3 3 0 0 0 3-3v-3h2v5a1 1 0 0 0 2 0v-5h2v3a3 3 0 0 0 3 3h3a1 1 0 0 0 0-2ZM6 14a2 2 0 0 1 0-4a1 1 0 0 0 1-1a5 5 0 0 1 9.73-1.61a1 1 0 0 0 .78.67A3 3 0 0 1 20 11a3 3 0 0 1-3 3Z"/>`),
		g.Group(children),
	)
}