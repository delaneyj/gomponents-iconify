package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Scooter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M16 17a2 2 0 1 0 4 0a2 2 0 1 0-4 0M4 17a2 2 0 1 0 4 0a2 2 0 1 0-4 0"/><path d="M8 17h5a6 6 0 0 1 5-5V7a2 2 0 0 0-2-2h-1"/></g>`),
		g.Group(children),
	)
}