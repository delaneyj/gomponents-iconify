package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CursorHoverThirtyTwoRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 10a4 4 0 0 1 4-4h20a4 4 0 0 1 4 4v10a4 4 0 0 1-2.328 3.635a2.996 2.996 0 0 0-.55-.756l-.892-.892A2 2 0 0 0 28 20V10a2 2 0 0 0-2-2H6a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h8v2H6a4 4 0 0 1-4-4V10Zm15.707 6.293A1 1 0 0 0 16 17v12a1 1 0 0 0 1.8.6l2.7-3.6H25a1 1 0 0 0 .707-1.707l-8-8ZM18 26v-6.586L22.586 24H20a1 1 0 0 0-.8.4L18 26Z"/>`),
		g.Group(children),
	)
}