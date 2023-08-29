package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrushOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 17a4 4 0 1 1 4 4H3v-4z"/><path d="M21 3a16 16 0 0 0-9.309 4.704M9.896 9.916A15.993 15.993 0 0 0 8.2 13.2M21 3a16 16 0 0 1-4.697 9.302m-2.195 1.786A15.993 15.993 0 0 1 10.8 15.8M3 3l18 18"/></g>`),
		g.Group(children),
	)
}