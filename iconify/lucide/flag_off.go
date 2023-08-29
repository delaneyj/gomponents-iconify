package lucide

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 2c3 0 5 2 8 2s4-1 4-1v11M4 22V4m0 11s1-1 4-1s5 2 8 2M2 2l20 20"/>`),
		g.Group(children),
	)
}