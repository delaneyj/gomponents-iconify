package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RadioButtonSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8 3a5 5 0 1 0 0 10A5 5 0 0 0 8 3ZM2 8a6 6 0 1 1 12 0A6 6 0 0 1 2 8Z"/>`),
		g.Group(children),
	)
}