package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AtomTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12a3 3 0 1 0 6 0a3 3 0 1 0-6 0m3 9v.01M3 9v.01M21 9v.01M8 20.1A9 9 0 0 1 3 13m13 7.1a9 9 0 0 0 5-7.1M6.2 5a9 9 0 0 1 11.4 0"/>`),
		g.Group(children),
	)
}