package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeadingHThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 9h6l-4 4h1a3 3 0 1 1-2.83 3.998M3 5v7m0 0v7m0-7h8m0-7v7m0 0v7"/>`),
		g.Group(children),
	)
}