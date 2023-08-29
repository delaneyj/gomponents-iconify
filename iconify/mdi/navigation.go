package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Navigation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2L4.5 20.29l.71.71L12 18l6.79 3l.71-.71L12 2Z"/>`),
		g.Group(children),
	)
}