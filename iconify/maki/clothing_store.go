package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClothingStore(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M4 2L1 4.5V7h3v6h7V7h3V4.5L11 2H9.5l-2 4l-2-4H4Z"/>`),
		g.Group(children),
	)
}