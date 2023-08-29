package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Blur(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 21a9.01 9.01 0 0 0 2.32-.302a9 9 0 0 0 1.74-16.733A9 9 0 1 0 12 21zm0-18v17m0-8h9m-9-3h8m-8-3h6m-6 12h6m-6-3h8"/>`),
		g.Group(children),
	)
}