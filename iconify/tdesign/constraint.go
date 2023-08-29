package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Constraint(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 8H6.2C3.83 8.026 2 9.851 2 12s1.83 3.975 4.2 4H10v2H6.186C2.809 17.967 0 15.338 0 12s2.81-5.967 6.186-6H10v2Zm4-2h3.75C21.155 6 24 8.641 24 12c0 3.36-2.845 6-6.25 6H14v-2h3.75c2.394 0 4.25-1.836 4.25-4c0-2.164-1.856-4-4.25-4H14V6Zm-7 5h10v2H7v-2Z"/>`),
		g.Group(children),
	)
}