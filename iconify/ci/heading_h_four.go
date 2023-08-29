package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeadingHFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m18 9l-2.5 8H20m0 0h1m-1 0v-3m0 3v2M3 5v7m0 0v7m0-7h8m0-7v7m0 0v7"/>`),
		g.Group(children),
	)
}