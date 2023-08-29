package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserRoundedBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><circle cx="12" cy="6" r="4"/><ellipse cx="12" cy="17" rx="7" ry="4"/></g>`),
		g.Group(children),
	)
}