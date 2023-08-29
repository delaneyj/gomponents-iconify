package typcn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17.568 9.432c0-2.55-.906-5.592-.944-5.72A1.001 1.001 0 0 0 15.666 3H8.334c-.441 0-.83.289-.958.712c-.038.128-.944 3.17-.944 5.72c0 2.735 1.984 5.011 4.587 5.477L11 15v4h-1a1 1 0 1 0 0 2h4a1 1 0 1 0 0-2h-1v-4l-.019-.092c2.603-.466 4.587-2.741 4.587-5.476zM12 13a3.565 3.565 0 0 1-3.511-3h7.021A3.563 3.563 0 0 1 12 13zM8.445 9c.062-1.468.422-3.093.653-4h5.803c.231.907.591 2.532.653 4H8.445z"/>`),
		g.Group(children),
	)
}