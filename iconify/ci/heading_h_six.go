package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeadingHSix(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15.402 14.525A2.974 2.974 0 0 0 16.5 18.6a3.01 3.01 0 0 0 4.099-1.092a2.973 2.973 0 0 0-1.099-4.075a3.01 3.01 0 0 0-4.098 1.092Zm0 0L19 8M3 5v7m0 0v7m0-7h8m0-7v7m0 0v7"/>`),
		g.Group(children),
	)
}