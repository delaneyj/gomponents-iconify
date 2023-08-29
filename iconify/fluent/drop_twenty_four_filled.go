package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DropTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M11.47 2.22a.75.75 0 0 1 1.06 0c.403.403 1.999 2.127 3.499 4.362C17.509 8.785 19 11.635 19 14.25c0 2.524-.746 4.479-2.044 5.806C15.659 21.38 13.889 22 12 22c-1.89 0-3.659-.619-4.956-1.944C5.746 18.729 5 16.774 5 14.25c0-2.615 1.492-5.465 2.971-7.668c1.5-2.235 3.096-3.96 3.499-4.362Z"/>`),
		g.Group(children),
	)
}