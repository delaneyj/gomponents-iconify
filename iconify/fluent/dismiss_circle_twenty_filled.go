package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DismissCircleTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 2a8 8 0 1 1 0 16a8 8 0 0 1 0-16ZM7.81 7.114a.5.5 0 0 0-.638.058l-.058.069a.5.5 0 0 0 .058.638L9.292 10l-2.12 2.121l-.058.07a.5.5 0 0 0 .058.637l.069.058a.5.5 0 0 0 .638-.058L10 10.708l2.121 2.12l.07.058a.5.5 0 0 0 .637-.058l.058-.069a.5.5 0 0 0-.058-.638L10.708 10l2.12-2.121l.058-.07a.5.5 0 0 0-.058-.637l-.069-.058a.5.5 0 0 0-.638.058L10 9.292l-2.121-2.12l-.07-.058Z"/>`),
		g.Group(children),
	)
}