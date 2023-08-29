package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TransitionLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 21a3 3 0 0 1-3-3V6a3 3 0 0 1 3-3m15 3v12a3 3 0 0 1-6 0V6a3 3 0 0 1 6 0zm-6 6H7m3-3l-3 3l3 3"/>`),
		g.Group(children),
	)
}