package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UsersGroupRoundedBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><circle cx="15" cy="6" r="3" opacity=".4"/><ellipse cx="16" cy="17" opacity=".4" rx="5" ry="3"/><circle cx="9.001" cy="6" r="4"/><ellipse cx="9.001" cy="17.001" rx="7" ry="4"/></g>`),
		g.Group(children),
	)
}