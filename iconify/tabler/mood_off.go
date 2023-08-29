package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoodOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M5.634 5.638a9 9 0 0 0 12.732 12.724m1.679-2.322A9 9 0 0 0 7.965 3.954M9 10h.01M15 10h.01"/><path d="M9.5 15a3.5 3.5 0 0 0 5 0M3 3l18 18"/></g>`),
		g.Group(children),
	)
}