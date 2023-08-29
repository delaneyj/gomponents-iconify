package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarTodoFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 3h4a1 1 0 0 1 1 1v16a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1h4V1h2v2h6V1h2v2ZM4 9v10h16V9H4Zm2 2h2v2H6v-2Zm0 4h2v2H6v-2Zm4-4h8v2h-8v-2Zm0 4h5v2h-5v-2Z"/>`),
		g.Group(children),
	)
}