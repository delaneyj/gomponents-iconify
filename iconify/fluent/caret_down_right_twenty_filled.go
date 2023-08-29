package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CaretDownRightTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M15 5.957c0-.89-1.077-1.337-1.707-.707L5.25 13.293c-.63.63-.184 1.707.707 1.707H13.5a1.5 1.5 0 0 0 1.5-1.5V5.957Z"/>`),
		g.Group(children),
	)
}