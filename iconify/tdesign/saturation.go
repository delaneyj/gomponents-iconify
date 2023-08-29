package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Saturation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 .07l7.068 7a9.856 9.856 0 0 1 0 14.029c-3.905 3.867-10.231 3.867-14.136 0a9.856 9.856 0 0 1 0-14.029L12 .07Zm0 2.814L6.34 8.491a7.856 7.856 0 0 0 0 11.187c3.125 3.095 8.195 3.095 11.32 0a7.856 7.856 0 0 0 0-11.187L12 2.884Zm-1 6.15h1a5.482 5.482 0 1 1 0 10.964h-1V9.035Zm2 2.147v6.671a3.483 3.483 0 0 0 0-6.671Z"/>`),
		g.Group(children),
	)
}