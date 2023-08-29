package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarTodoLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 1v2h6V1h2v2h4a1 1 0 0 1 1 1v16a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1h4V1h2Zm11 10H4v8h16v-8ZM8 14v2H6v-2h2Zm10 0v2h-8v-2h8ZM7 5H4v4h16V5h-3v2h-2V5H9v2H7V5Z"/>`),
		g.Group(children),
	)
}