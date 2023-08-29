package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pinwheel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 12c0-3 2.5-5.5 5.5-5.5S23 9 23 12H12m0 0c0 3-2.5 5.5-5.5 5.5S1 15 1 12h11m0 0c-3 0-5.5-2.5-5.5-5.5S9 1 12 1v11m0 0c3 0 5.5 2.5 5.5 5.5S15 23 12 23V12Z"/>`),
		g.Group(children),
	)
}