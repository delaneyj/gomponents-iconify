package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MathOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m14 19l2.5-2.5m2-2L20 13M3 3l18 18M19 5h-7l-.646 2.262m-.906 3.169L8 19l-3-6H3"/>`),
		g.Group(children),
	)
}