package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartDots(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 3v18h18"/><path d="M7 9a2 2 0 1 0 4 0a2 2 0 1 0-4 0m10-2a2 2 0 1 0 4 0a2 2 0 1 0-4 0m-5 8a2 2 0 1 0 4 0a2 2 0 1 0-4 0m-1.84-4.38l2.34 2.88m2.588-.172l2.837-4.586"/></g>`),
		g.Group(children),
	)
}