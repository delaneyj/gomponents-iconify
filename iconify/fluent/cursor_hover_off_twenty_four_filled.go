package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CursorHoverOffTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3.28 2.22a.75.75 0 1 0-1.06 1.06l1.045 1.046A2.5 2.5 0 0 0 2 6.5v9A2.5 2.5 0 0 0 4.5 18H9v-7.25c0-.204.035-.402.102-.588l.898.899V21.25a.75.75 0 0 0 1.368.425l2.467-3.588l4.042.85l2.842 2.843a.75.75 0 0 0 1.061-1.06L3.28 2.22ZM7.182 4L20.81 17.63A2.499 2.499 0 0 0 22 15.5v-9A2.5 2.5 0 0 0 19.5 4H7.182Z"/>`),
		g.Group(children),
	)
}