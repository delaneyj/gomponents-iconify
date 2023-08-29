package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilmstripSplitSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8.5 1.5a.5.5 0 0 0-1 0v13a.5.5 0 0 0 1 0v-13ZM1 5.5A2.5 2.5 0 0 1 3.5 3h3v10h-3A2.5 2.5 0 0 1 1 10.5v-5Zm2 0v1a.5.5 0 0 0 1 0v-1a.5.5 0 0 0-1 0ZM3.5 9a.5.5 0 0 0-.5.5v1a.5.5 0 0 0 1 0v-1a.5.5 0 0 0-.5-.5Zm6 4V3h3A2.5 2.5 0 0 1 15 5.5v5a2.5 2.5 0 0 1-2.5 2.5h-3ZM12 5.5v1a.5.5 0 0 0 1 0v-1a.5.5 0 0 0-1 0Zm.5 3.5a.5.5 0 0 0-.5.5v1a.5.5 0 0 0 1 0v-1a.5.5 0 0 0-.5-.5Z"/>`),
		g.Group(children),
	)
}