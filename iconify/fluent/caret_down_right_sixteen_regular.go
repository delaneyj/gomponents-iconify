package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CaretDownRightSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10.293 4.25c.63-.63 1.707-.184 1.707.707V10.5a1.5 1.5 0 0 1-1.5 1.5H4.957c-.89 0-1.337-1.077-.707-1.707l6.043-6.043Zm.707.707L4.957 11H10.5a.5.5 0 0 0 .5-.5V4.957Z"/>`),
		g.Group(children),
	)
}