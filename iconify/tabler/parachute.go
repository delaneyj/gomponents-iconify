package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Parachute(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M22 12a10 10 0 1 0-20 0"/><path d="M22 12c0-1.66-1.46-3-3.25-3c-1.8 0-3.25 1.34-3.25 3c0-1.66-1.57-3-3.5-3s-3.5 1.34-3.5 3c0-1.66-1.46-3-3.25-3C3.45 9 2 10.34 2 12m0 0l10 10l-3.5-10m7 0L12 22l10-10"/></g>`),
		g.Group(children),
	)
}