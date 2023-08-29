package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LuggageThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 2a1 1 0 0 1 1-1h10a1 1 0 1 1 0 2h-.5v2h2A4.5 4.5 0 0 1 27 9.5v15a4.5 4.5 0 0 1-4 4.473V30a1 1 0 1 1-2 0v-1H11v1a1 1 0 1 1-2 0v-1.027A4.5 4.5 0 0 1 5 24.5v-15A4.5 4.5 0 0 1 9.5 5H12V3h-1a1 1 0 0 1-1-1Zm4 1v2h4.5V3H14Zm-3 8a1 1 0 1 0 0 2h10a1 1 0 1 0 0-2H11Z"/>`),
		g.Group(children),
	)
}