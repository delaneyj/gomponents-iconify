package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrainCarBox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 7v10h1a2 2 0 1 0 4 0h12a2 2 0 1 0 4 0h1V7H1Z"/>`),
		g.Group(children),
	)
}