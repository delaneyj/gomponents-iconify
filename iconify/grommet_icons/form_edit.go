package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormEdit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="m11.996 8.336l3.497 3.498l-3.497-3.498Zm5.54-.54a.994.994 0 0 1-.004 1.416l-7.451 7.451L6 17.83l1.166-4.08l7.451-7.452a.997.997 0 0 1 1.416-.005l1.504 1.504Z"/>`),
		g.Group(children),
	)
}