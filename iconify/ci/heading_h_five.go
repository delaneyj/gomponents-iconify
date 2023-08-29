package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeadingHFive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 9h-4l-1.25 5.016a2.998 2.998 0 0 1 3.555-.717a3 3 0 1 1-3.555 4.685M3 5v7m0 0v7m0-7h8m0-7v7m0 0v7"/>`),
		g.Group(children),
	)
}