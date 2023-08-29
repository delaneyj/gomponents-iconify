package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeartHalfOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 8.5c0 2.7 2.75 5.37 7 9.24V7.2C10.42 5.91 9 5 7.5 5C5.5 5 4 6.5 4 8.5m9-1.3v13.24l-1 .91l-1.45-1.32C5.4 15.36 2 12.27 2 8.5C2 5.41 4.42 3 7.5 3C10 3 13 5 13 7.2Z"/>`),
		g.Group(children),
	)
}