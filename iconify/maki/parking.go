package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Parking(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M4 2v11h2V9h2.5a3.5 3.5 0 1 0 0-7H4Zm2 5V4h2.5a1.5 1.5 0 1 1 0 3H6Z"/>`),
		g.Group(children),
	)
}