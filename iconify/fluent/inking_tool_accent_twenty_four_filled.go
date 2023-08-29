package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InkingToolAccentTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4 7h16.25a.25.25 0 0 0 .243-.193l.007-.057V3H3.75v3.75a.25.25 0 0 0 .193.243L4 7Zm7.976 6.946l2.641-5.947l.883-.499l.758.499l-2.911 6.556a.75.75 0 1 1-1.371-.61ZM13 19c0 1.105-.448 2-1 2s-1-.895-1-2s.448-2 1-2s1 .895 1 2Z"/>`),
		g.Group(children),
	)
}