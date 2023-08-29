package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SumOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 18a1 1 0 0 1-1 1H6l6-7M9 5h8a1 1 0 0 1 1 1v2M3 3l18 18"/>`),
		g.Group(children),
	)
}