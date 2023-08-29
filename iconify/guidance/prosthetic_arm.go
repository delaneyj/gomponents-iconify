package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProstheticArm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M10 18v6m4 0v-6.5a2 2 0 0 1 2-2v-8s-1.5-1-4-1s-4 1-4 1V11m1.5 4.5H8m0 0H6.5m1.5 0V12m3.85-7.5s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0c0 1.25-1.596 2.25-1.596 2.25h-.3Z"/>`),
		g.Group(children),
	)
}