package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SlideTextFortyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4 12.75A4.75 4.75 0 0 1 8.75 8h30.5A4.75 4.75 0 0 1 44 12.75v22.5A4.75 4.75 0 0 1 39.25 40H8.75A4.75 4.75 0 0 1 4 35.25v-22.5Zm9.25 9.75a1.25 1.25 0 1 0 0 2.5h17.5a1.25 1.25 0 1 0 0-2.5h-17.5ZM12 30.25c0 .69.56 1.25 1.25 1.25h13.5a1.25 1.25 0 1 0 0-2.5h-13.5c-.69 0-1.25.56-1.25 1.25ZM13.25 16a1.25 1.25 0 1 0 0 2.5h9.5a1.25 1.25 0 1 0 0-2.5h-9.5Z"/>`),
		g.Group(children),
	)
}