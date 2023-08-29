package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScissorsOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4.432 4.442a3 3 0 1 0 4.114 4.146M3 17a3 3 0 1 0 6 0a3 3 0 1 0-6 0m5.6-1.6L12 12m2-2l5-5M3 3l18 18"/>`),
		g.Group(children),
	)
}