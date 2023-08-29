package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CaretDownRightSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12 4.957c0-.89-1.077-1.337-1.707-.707L4.25 10.293c-.63.63-.184 1.707.707 1.707H10.5a1.5 1.5 0 0 0 1.5-1.5V4.957Z"/>`),
		g.Group(children),
	)
}