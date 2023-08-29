package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UsersGroupRoundedLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><circle cx="9" cy="6" r="4"/><path stroke-linecap="round" d="M15 9a3 3 0 1 0 0-6"/><ellipse cx="9" cy="17" rx="7" ry="4"/><path stroke-linecap="round" d="M18 14c1.754.385 3 1.359 3 2.5c0 1.03-1.014 1.923-2.5 2.37"/></g>`),
		g.Group(children),
	)
}