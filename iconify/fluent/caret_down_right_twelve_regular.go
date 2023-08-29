package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CaretDownRightTwelveRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7.298 3.256c.63-.63 1.707-.184 1.707.707v3.543a1.5 1.5 0 0 1-1.5 1.5H3.962c-.89 0-1.337-1.077-.707-1.707l4.043-4.043Zm.707.707L3.962 8.006h3.543a.5.5 0 0 0 .5-.5V3.963Z"/>`),
		g.Group(children),
	)
}