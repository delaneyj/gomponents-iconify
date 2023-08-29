package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EqualCircleTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M22 12c0-5.523-4.477-10-10-10S2 6.477 2 12s4.477 10 10 10s10-4.477 10-10Zm-5.748-1.5h-8.5a.75.75 0 0 1 0-1.5h8.5a.75.75 0 0 1 0 1.5Zm0 4.5h-8.5a.75.75 0 0 1 0-1.5h8.5a.75.75 0 0 1 0 1.5Z"/>`),
		g.Group(children),
	)
}