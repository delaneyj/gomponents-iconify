package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BracesOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5.176 5.177C5.063 5.428 5 5.707 5 6v3c0 1.657-.895 3-2 3c1.105 0 2 1.343 2 3v3a2 2 0 0 0 2 2M17 4a2 2 0 0 1 2 2v3c0 1.657.895 3 2 3c-1.105 0-2 1.343-2 3m-.176 3.821A2 2 0 0 1 17 20M3 3l18 18"/>`),
		g.Group(children),
	)
}