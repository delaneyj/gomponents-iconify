package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Book(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 6a4 4 0 0 1 4-4h14v20H7a4 4 0 0 1-4-4V6Zm2 8.535A3.982 3.982 0 0 1 7 14h12V4H7a2 2 0 0 0-2 2v8.535ZM19 16H7a2 2 0 1 0 0 4h12v-4ZM10 6h7v2h-7V6Z"/>`),
		g.Group(children),
	)
}