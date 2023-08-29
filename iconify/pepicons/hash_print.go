package pepicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HashPrint(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round"><path stroke-width="3" d="M5.5 8.75h11m-11 5.5h11M8.75 5l-1 12m6.5-12l-1 12" opacity=".8"/><path d="M3.5 7.75h13m-13 5.5h13M7.75 4l-1 12m6.5-12l-1 12"/></g>`),
		g.Group(children),
	)
}