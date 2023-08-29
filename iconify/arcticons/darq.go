package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Darq(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.54 2.777A21.502 21.502 0 1 1 24 45.5a21.647 21.647 0 0 1-3.48-.28"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.52 45.22a21.504 21.504 0 0 1 .02-42.443m.02 42.283V2.848m3.986 4.651c8.377 0 16.928 6.173 16.928 16.501s-8.551 16.501-16.928 16.501m0 0V7.5"/>`),
		g.Group(children),
	)
}