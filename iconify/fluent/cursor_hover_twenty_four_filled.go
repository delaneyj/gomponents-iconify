package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CursorHoverTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4.5 4A2.5 2.5 0 0 0 2 6.5v9A2.5 2.5 0 0 0 4.5 18H9v-7.25a1.75 1.75 0 0 1 2.785-1.411l7.5 7.5c.375.275.615.679.69 1.116A2.5 2.5 0 0 0 22 15.5v-9A2.5 2.5 0 0 0 19.5 4h-15Zm6.78 6.22a.75.75 0 0 0-1.28.53v10.5a.75.75 0 0 0 1.368.425l2.467-3.588l4.26.897a.75.75 0 0 0 .685-1.264l-7.5-7.5Z"/>`),
		g.Group(children),
	)
}