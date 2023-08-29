package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoodShare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M20.942 13.018A9 9 0 1 0 12 21M9 10h.01M15 10h.01"/><path d="M9.5 15c.658.672 1.56 1 2.5 1c.213 0 .424-.017.63-.05M16 22l5-5m0 4.5V17h-4.5"/></g>`),
		g.Group(children),
	)
}