package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EraserOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m3 3l18 18m-2-1H8.5l-4.21-4.3a1 1 0 0 1 0-1.41l5-4.993m2.009-2.01l3-3a1 1 0 0 1 1.41 0l5 5a1 1 0 0 1 0 1.41c-1.417 1.431-2.406 2.432-2.97 3m-2.02 2.043l-4.211 4.256M18 13.3L11.7 7"/>`),
		g.Group(children),
	)
}