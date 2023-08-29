package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Commandline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M2.02 9.748v-.002v.002zm.002-.002l5.77 5.773l-5.77 5.77l2.12 2.123l7.895-7.895l-7.894-7.895l-2.12 2.123zM12.248 23.27h14.42v-3h-14.42v3zm4.335-6.25h10.084v-3H16.583v3zm-4.335-9.25v3h14.42v-3h-14.42z"/>`),
		g.Group(children),
	)
}