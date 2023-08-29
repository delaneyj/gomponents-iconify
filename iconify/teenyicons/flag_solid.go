package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M0 .5A.5.5 0 0 1 .5 0h14a.5.5 0 0 1 .457.703L13.047 5l1.91 4.297A.5.5 0 0 1 14.5 10H1v5H0V.5Z"/>`),
		g.Group(children),
	)
}