package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HighlightAccentTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M20.003 4L20 7.75l-.007.102a.754.754 0 0 1-.71.648H4.737l-.089-.006A.75.75 0 0 1 4 7.751V4h16.003Zm-4.496 9.5H8.5V21l6.576-3.106a.75.75 0 0 0 .424-.572l.008-.107l-.001-3.715Z"/>`),
		g.Group(children),
	)
}