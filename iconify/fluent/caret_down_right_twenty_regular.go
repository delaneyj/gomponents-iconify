package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CaretDownRightTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M13.293 5.25c.63-.63 1.707-.184 1.707.707V13.5a1.5 1.5 0 0 1-1.5 1.5H5.957c-.89 0-1.337-1.077-.707-1.707l8.043-8.043Zm.707.707L5.957 14H13.5a.5.5 0 0 0 .5-.5V5.957Z"/>`),
		g.Group(children),
	)
}