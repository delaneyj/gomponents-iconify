package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataSharing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.52 3.87a5 5 0 0 0-9.08.13H7a3 3 0 0 0-3 3v4a1 1 0 0 0 2 0V7a1 1 0 0 1 1-1h2.78A3 3 0 0 0 9 8a3 3 0 0 0 3 3h7.33a3.66 3.66 0 0 0 1.19-7.13ZM19.33 9H12a1 1 0 0 1 0-2a1 1 0 0 0 1-1a3 3 0 0 1 5.84-1a1 1 0 0 0 .78.67A1.65 1.65 0 0 1 21 7.33A1.67 1.67 0 0 1 19.33 9ZM19 13a1 1 0 0 0-1 1v3a1 1 0 0 1-1 1h-2.26a3.66 3.66 0 0 0-2.22-2.13a5 5 0 0 0-9.45 1.28A3 3 0 0 0 4 23h7.33a3.66 3.66 0 0 0 3.6-3H17a3 3 0 0 0 3-3v-3a1 1 0 0 0-1-1Zm-7.67 8H4a1 1 0 0 1 0-2a1 1 0 0 0 1-1a3 3 0 0 1 5.84-1a1 1 0 0 0 .78.67A1.65 1.65 0 0 1 13 19.33A1.67 1.67 0 0 1 11.33 21Z"/>`),
		g.Group(children),
	)
}