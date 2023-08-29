package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Connect(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M10 21c-2.5 2.5-5 2-7 0s-2.5-4.5 0-7l3-3l7 7l-3 3Zm4-18c2.5-2.5 5-2 7.001 0c2.001 2 2.499 4.5 0 7l-3 3L11 6l3-3Zm-3 7l-2.5 2.5L11 10Zm3 3l-2.5 2.5L14 13Z"/>`),
		g.Group(children),
	)
}