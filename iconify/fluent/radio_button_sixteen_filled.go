package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RadioButtonSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8 11a3 3 0 1 0 0-6a3 3 0 0 0 0 6Zm0-9a6 6 0 1 0 0 12A6 6 0 0 0 8 2ZM3 8a5 5 0 1 1 10 0A5 5 0 0 1 3 8Z"/>`),
		g.Group(children),
	)
}