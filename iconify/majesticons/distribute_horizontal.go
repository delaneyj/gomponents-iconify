package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DistributeHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M5 20V4m14 16V4"/><rect width="10" height="4" x="10" y="17" fill="currentColor" rx="2" transform="rotate(-90 10 17)"/></g>`),
		g.Group(children),
	)
}