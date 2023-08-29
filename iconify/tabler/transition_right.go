package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TransitionRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 3a3 3 0 0 1 3 3v12a3 3 0 0 1-3 3M3 18V6a3 3 0 1 1 6 0v12a3 3 0 0 1-6 0zm6-6h8m-3 3l3-3l-3-3"/>`),
		g.Group(children),
	)
}