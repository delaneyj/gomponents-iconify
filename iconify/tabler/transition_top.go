package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TransitionTop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 6a3 3 0 0 0-3-3H6a3 3 0 0 0-3 3m3 15h12a3 3 0 0 0 0-6H6a3 3 0 0 0 0 6zm6-6V7m-3 3l3-3l3 3"/>`),
		g.Group(children),
	)
}