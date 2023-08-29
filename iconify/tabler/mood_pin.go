package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoodPin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M21 12a9 9 0 1 0-8.352 8.977M9 10h.01M15 10h.01"/><path d="M9.5 15c.658.672 1.56 1 2.5 1c.102 0 .203-.004.304-.012m8.817 4.133a3 3 0 1 0-4.242 0c.418.419 1.125 1.045 2.121 1.879c1.051-.89 1.759-1.516 2.121-1.879zM19 18v.01"/></g>`),
		g.Group(children),
	)
}